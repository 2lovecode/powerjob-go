package engine

type Remote interface {
	Start(config *Config) (output *Output)
	Close() error
}
