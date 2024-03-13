package common

type Config struct {
	appName                    string
	port                       int
	serverAddress              []string
	protocol                   int //
	maxResultLength            int
	userContext                interface{}
	storeStrategy              interface{}
	allowLazyConnectServer     bool
	maxAppendedWfContextLength int
	systemMetricsCollector     interface{}
	processorFactoryList       []interface{}
	tag                        string
	maxLightweightTaskNum      int
	maxHeavyweightTaskNum      int
	healthReportInterval       int
}

func NewConfig() *Config {
	return &Config{
		appName:                    "",
		port:                       27777,
		serverAddress:              nil,
		protocol:                   0,
		maxResultLength:            8096,
		userContext:                nil,
		storeStrategy:              nil,
		allowLazyConnectServer:     false,
		maxAppendedWfContextLength: 8192,
		systemMetricsCollector:     nil,
		processorFactoryList:       nil,
		tag:                        "",
		maxLightweightTaskNum:      1024,
		maxHeavyweightTaskNum:      64,
		healthReportInterval:       10,
	}
}
