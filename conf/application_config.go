package conf

type Configuration struct {
	App        App
	AuthType   string `env:"AUTH_TYPE, default=no_auth"`
	BasicAuth  BasicAuth
	Db         Db
	Connect    Connect
	Kafka      Kafka
	Connectors []Connector
}

type BasicAuth struct {
	Username string `env:"BASIC_AUTH_USERNAME"`
	Password string `env:"BASIC_AUTH_PASSWORD"`
}

type App struct {
	Host string `env:"APP_HOST, default=localhost"`
	Port string `env:"APP_PORT, default=8080"`
	Log  struct {
		Level string `env:"APP_LOG_LEVEL, default=warn"`
	}
}

type Db struct {
	Host     string `env:"DB_HOST, required=true"`
	Port     string `env:"DB_PORT, required=true"`
	Name     string `env:"DB_NAME, required=true"`
	Username string `env:"DB_USERNAME, required=true"`
	Password string `env:"DB_PASSWORD, required=true"`
}

type Connect struct {
	Url string
}

type Kafka struct {
	BootstrapServers string `mapstructure:"bootstrap-servers"`
}

type Connector struct {
	ID             int
	Name           string
	TableName      string `mapstructure:"table-name"`
	Topic          string
	Type           string
	KeepSchema     bool `mapstructure:"keep-schema"`
	KeepTombstones bool `mapstructure:"keep-tombstones"`
}
