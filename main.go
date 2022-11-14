package main

import (
	"flag"
	"fmt"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/globals"
	"github.com/gunjdesai/kafka-connect-connector-manager/handlers"
	"github.com/gunjdesai/kafka-connect-connector-manager/helpers"
	_ "go.uber.org/automaxprocs"
	"os"
)

func main() {

	// Load Global variables
	globals.Bootstrap()

	// Invoking CLI variables
	mode, connectorName, skipValidation := readFromCli()

	if !skipValidation {
		//TODO: call validation directly
	}

	if connectorName == "" && mode != constants.STATUS_MODE && mode != constants.RESUME_ALL_MODE && mode != constants.PAUSE_ALL_MODE {
		helpers.PrintLog("Connector Name cannot be empty, pass connector name by using -connector-name or --connector-name", "", false, nil)
		os.Exit(10)
	}

	switch mode {

	case constants.UPSERT_MODE:
		handlers.PutConnector(connectorName)
	case constants.DELETE_MODE:
		handlers.DeleteConnector(connectorName)
	case constants.STATUS_MODE:
		handlers.StatusConnector()
	case constants.PAUSE_ALL_MODE:
		handlers.PauseAllConnectors()
	case constants.RESUME_ALL_MODE:
		handlers.ResumeAllConnectors()
	default:
		helpers.PrintLog("Invalid Mode, Shutting down", "", false, nil)
		os.Exit(10)

	}

}

func readFromCli() (string, string, bool) {

	modeOptions := fmt.Sprintf("Options are %s, %s, %s, %s & %s", constants.UPSERT_MODE, constants.DELETE_MODE, constants.STATUS_MODE, constants.PAUSE_ALL_MODE, constants.RESUME_ALL_MODE)
	mode := flag.String("mode", "", "Mode to run the utility under, "+modeOptions)
	connectorName := flag.String("connector-name", "", "Name of the connector you want to modify")
	skipValidation := flag.Bool("skip-validation", false, "Skip validation of config, default value is false")
	flag.Parse()

	if *mode != constants.UPSERT_MODE && *mode != constants.DELETE_MODE && *mode != constants.STATUS_MODE && *mode != constants.PAUSE_ALL_MODE && *mode != constants.RESUME_ALL_MODE {
		helpers.PrintLog("Invalid Mode, ", modeOptions, false, nil)
		os.Exit(10)
	}

	return *mode, *connectorName, *skipValidation

}
