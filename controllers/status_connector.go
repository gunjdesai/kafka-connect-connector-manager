package controllers

import (
	"encoding/json"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/data/response"
	"github.com/gunjdesai/kafka-connect-connector-manager/globals"
	"github.com/gunjdesai/kafka-connect-connector-manager/helpers"
	"github.com/gunjdesai/kafka-connect-connector-manager/network"
	"strconv"
	"time"
)

func getConnectorStatus(baseUrl string, name string) {

	url := baseUrl + constants.SLASH_SEPERATOR_URL + name + constants.SLASH_SEPERATOR_URL + constants.STATUS_URL
	var body response.ConnectorStatus
	var errBody response.ApiErrorResponse
	var err error

	resp, err := network.MakeHttpRequest(constants.GET_METHOD, url, nil)
	if err != nil {
		helpers.PrintLog("Status Check Failed for ", name, false, nil)
		helpers.PrintLog("Error encountered while getting status:", err.Error(), false, nil)
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&body)
		if err != nil {
			helpers.PrintLog("Issue with parsing Connectors status api data:", err.Error(), false, body)
		}
		helpers.PrintConnectorStatusLog(body.Name, body.Connector, body.Tasks)
	} else {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&errBody)
		helpers.PrintLog("Status Check Failed for ", name, false, nil)
		helpers.PrintLog("Issue with decoding Connector Response:", err.Error(), false, errBody)
		return
	}

}

func GetAllConnectorStatus() {

	url := globals.Config.Connect.Url + constants.SLASH_SEPERATOR_URL + constants.CONNECTORS_URL
	var body response.AllConnectors
	var errBody response.ApiErrorResponse
	var err error
	resp, err := network.MakeHttpRequest(constants.GET_METHOD, url, nil)
	if err != nil {
		helpers.PrintLog("Error encountered while getting connectors:", err.Error(), false, nil)
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&body)
		if err != nil {
			helpers.PrintLog("Issue with parsing Connectors api data:", err.Error(), false, body)
		}
		helpers.PrintLog("Fetched Total Connectors:", strconv.Itoa(len(body)), true, body)

		for _, c := range body {
			time.Sleep(1000 * time.Millisecond)
			getConnectorStatus(url, c)
		}
	} else {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&errBody)
		helpers.PrintLog("Issue with fetching Connectors:", err.Error(), false, errBody)
		return
	}

}
