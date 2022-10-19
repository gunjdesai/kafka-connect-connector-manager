package globals

import (
	"github.com/gunjdesai/kafka-connect-connector-manager/conf"
	"github.com/gunjdesai/kafka-connect-connector-manager/constants"
	"github.com/gunjdesai/kafka-connect-connector-manager/helpers"
	"github.com/gunjdesai/kafka-connect-connector-manager/loggers"
	"os"
)

var (
	Config *conf.Configuration
	Log    *loggers.Logger
)

func Bootstrap() {

	var err error
	var confFile = helpers.GetEnv() + constants.DOT_SEPERATOR_FILE + constants.TYPE_FILE
	Config, err = conf.New(constants.CONFIG_PATH_FILE, confFile)

	if err != nil {
		helpers.PrintLog("Error on bootstraping config:", err.Error(), false, nil)
		os.Exit(10)
	}

	Log, err = loggers.New(Config.App.Log.Level)

	if err != nil {
		helpers.PrintLog("Error on bootstraping logger:", err.Error(), false, nil)
		os.Exit(10)
	}

}
