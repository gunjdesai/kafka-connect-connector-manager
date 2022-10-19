package conf

import (
	"errors"
	"github.com/Netflix/go-env"
	"github.com/gunjdesai/kafka-connect-connector-manager/helpers"
	"github.com/spf13/viper"
	"path"
)

var Config *Configuration

func New(location string, fileName string) (*Configuration, error) {

	Config = &Configuration{}
	filePath := path.Join(helpers.GetRootDir(), location, fileName)
	viper.SetConfigFile(filePath)

	if err := viper.ReadInConfig(); err != nil {
		err = errors.New("Error reading config file, " + err.Error())
		return nil, err
	}

	if err := viper.Unmarshal(Config); err != nil {
		err = errors.New("Unable to decode to struct, " + err.Error())
		return nil, err
	}

	if _, err := env.UnmarshalFromEnviron(Config); err != nil {
		err = errors.New("Environment variable override failure, " + err.Error())
		return nil, err
	}

	return Config, nil

}
