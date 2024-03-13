package engine

type PowerJobRemoteEngine struct {
}

func NewPowerJobRemoteEngine() *PowerJobRemoteEngine {
	return &PowerJobRemoteEngine{}
}

func (eg *PowerJobRemoteEngine) Start(config *Config) (output *Output) {

	return
}

func (eg *PowerJobRemoteEngine) Close() error {
	return nil
}
