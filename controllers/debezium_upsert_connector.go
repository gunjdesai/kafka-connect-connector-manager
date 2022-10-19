package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gunjdesai/kafka-connect-connector-manager/conf"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/data/response"
	"github.com/gunjdesai/kafka-connect-connector-manager/globals"
	"github.com/gunjdesai/kafka-connect-connector-manager/helpers"
	"github.com/gunjdesai/kafka-connect-connector-manager/network"
	"strings"
)

func PutDebeziumConnector(connector conf.Connector) {

	var body response.ApiErrorResponse
	var err error
	url := globals.Config.Connect.Url + constants.SLASH_SEPERATOR_URL + constants.CONNECTORS_URL +
		constants.SLASH_SEPERATOR_URL + connector.Name + constants.SLASH_SEPERATOR_URL + constants.CONFIG_URL
	requestBody, err := json.Marshal(addConnectorConfig(connector))
	if err != nil {
		helpers.PrintLog("Error encountered while attempting to creating connector config:", err.Error(), false, nil)
	}

	resp, err := network.MakeHttpRequest(constants.PUT_METHOD, url, bytes.NewBuffer(requestBody))
	if err != nil {
		helpers.PrintLog("Error encountered while attempting to create connector:", err.Error(), false, nil)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		helpers.PrintLog("Successfully created Connector:", connector.Name, true, nil)
	} else {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&body)
		helpers.PrintLog("Issue with Creating Connector:", connector.Name, false, body)
	}

}

func addConnectorConfig(connector conf.Connector) map[string]interface{} {

	config := make(map[string]interface{})
	transforms := []string{}

	config["database.server.name"] = connector.Name
	config["max.retries"] = 5
	config["database.server.id"] = connector.ID
	config["database.hostname"] = globals.Config.Db.Host
	config["database.user"] = globals.Config.Db.Username
	config["database.password"] = globals.Config.Db.Password
	config["database.history.kafka.bootstrap.servers"] = globals.Config.Kafka.BootstrapServers
	config["database.history.kafka.topic"] = "dbz.tables.history." + connector.TableName
	config["database.history.skip.unparseable.ddl"] = true
	config["database.include.list"] = globals.Config.Db.Name
	config["table.include.list"] = connector.TableName
	config["snapshot.mode"] = "schema_only"
	config["snapshot.locking.mode"] = "none"
	config["snapshot.delay.ms"] = 0
	config["include.query"] = false
	config["binlog.buffer.size"] = 0
	config["decimal.handling.mode"] = "double"

	if !connector.KeepSchema {
		config["value.converter"] = "org.apache.kafka.connect.json.JsonConverter"
		config["value.converter.schemas.enable"] = "false"
		transforms = append(transforms, "unwrap")
	}

	if !connector.KeepTombstones {
		config["transforms.unwrap.type"] = "io.debezium.transforms.ExtractNewRecordState"
		config["transforms.unwrap.drop.tombstones"] = true
		config["transforms.unwrap.delete.handling.mode"] = "rewrite"
		transforms = append(transforms, "Reroute")
	}

	//TODO: Add DB Type here
	config["connector.class"] = "io.debezium.connector.mysql.MySqlConnector"

	if connector.Topic != "" {
		config["transforms.Reroute.type"] = "io.debezium.transforms.ByLogicalTableRouter"
		config["transforms.Reroute.topic.regex"] = connector.Name + "." + connector.TableName
		config["transforms.Reroute.topic.replacement"] = connector.Topic
	}

	config["transforms"] = strings.Join(transforms[:], ",")

	return config

}
