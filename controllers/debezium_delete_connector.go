package controllers

import (
	"encoding/json"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/data/response"
	"github.com/gunjdesai/kafka-connect-connector-manager/globals"
	"github.com/gunjdesai/kafka-connect-connector-manager/helpers"
	"github.com/gunjdesai/kafka-connect-connector-manager/network"
)

func DeleteDebeziumConnector(name string) {

	url := globals.Config.Connect.Url + constants.SLASH_SEPERATOR_URL + constants.CONNECTORS_URL + constants.SLASH_SEPERATOR_URL + name
	var body response.ApiErrorResponse
	var err error
	resp, err := network.MakeHttpRequest(constants.DELETE_METHOD, url, nil)
	if err != nil {
		helpers.PrintLog("Error encountered while attempting delete:", err.Error(), false, nil)
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		helpers.PrintLog("Successfully deleted Connector:", name, true, nil)
	} else {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&body)
		helpers.PrintLog("Issue with Deleting Connector:", name, false, body)
	}

}
