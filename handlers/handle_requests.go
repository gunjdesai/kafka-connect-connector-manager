package handlers

import (
	"errors"
	"github.com/gunjdesai/kafka-connect-connector-manager/conf"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/controllers"
	"github.com/gunjdesai/kafka-connect-connector-manager/globals"
	"github.com/gunjdesai/kafka-connect-connector-manager/helpers"
	"os"
)

func PutConnector(name string) {

	config, err := getConnectorConfig(name)

	if err != nil {
		helpers.PrintLog("Shutting down..", err.Error(), false, nil)
		os.Exit(10)
	}

	switch config.Type {

	case constants.DEBEZIUM_CONNECTOR_TYPE:
		controllers.PutDebeziumConnector(*config)
	default:
		helpers.PrintLog("Invalid Connector Type, Current list of supported connectors are:", constants.DEBEZIUM_CONNECTOR_TYPE, false, nil)
		os.Exit(10)

	}

}

func DeleteConnector(name string) {
	controllers.DeleteDebeziumConnector(name)
}

func StatusConnector() {
	controllers.GetAllConnectorStatus()
}

func getConnectorConfig(name string) (*conf.Connector, error) {

	for _, v := range globals.Config.Connectors {
		if v.Name == name {
			return &v, nil
		}
	}

	err := errors.New("Invalid Connector Name: No connector found of the name " + name)
	return nil, err

}
