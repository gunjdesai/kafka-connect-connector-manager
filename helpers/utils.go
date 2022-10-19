package helpers

import (
	"fmt"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/data/response"
	"github.com/k0kubun/pp"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func GetEnv() string {

	env := os.Getenv(constants.VARIABLE_ENV)
	envs := []string{constants.CLASSZOO1_ENV, constants.SUITECRM_ENV, constants.DAILERDB_ENV}
	for _, v := range envs {
		if v == env {
			println("Running Env: " + v)
			return env
		}
	}

	println("Running Default Env: " + constants.CLASSZOO1_ENV)
	return constants.CLASSZOO1_ENV

}

func getConnectorStatusColor(state string) string {

	if state == "RUNNING" {
		return constants.GREEN_COLOR
	} else if state == "FAILED" {
		return constants.RED_COLOR
	} else {
		return constants.YELLOW_COLOR
	}

}

func PrintConnectorStatusLog(name string, connector response.Connector, tasks []response.Task) {

	var working = true
	print(constants.GREEN_COLOR, "Connector: ")
	print(constants.ORANGE_COLOR, name)
	print(constants.YELLOW_COLOR, " || ")
	print(getConnectorStatusColor(connector.State), " Status: "+connector.State)
	if (connector.State) == "FAILED" {
		working = false
	}
	print(constants.YELLOW_COLOR, " || ")
	for k, t := range tasks {
		print(constants.BLUE_COLOR, " ", t.ID)
		print(constants.CYAN_COLOR, ": ")
		print(getConnectorStatusColor(t.State), t.State)

		if k != len(tasks)-1 {
			print(constants.YELLOW_COLOR, " || ")
		}

		if (t.State) == "FAILED" {
			working = false
		}

	}
	println("")
	if !working {
		println(constants.RED_COLOR, "===================================================================================================================")
	}

}

func PrintLog(text string, name string, success bool, data interface{}) {

	var textColor = constants.RED_COLOR
	var nameColor = constants.YELLOW_COLOR

	if success {
		textColor = constants.GREEN_COLOR
		nameColor = constants.WHITE_COLOR
	}

	fmt.Print(textColor, text)
	if name != "" {
		fmt.Println(nameColor, name)
	}
	if data != nil {
		pp.Println(data)
	}

}
