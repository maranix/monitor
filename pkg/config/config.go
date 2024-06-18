package config

type Config interface {
	GetDebounce() float32
	GetIgnoreTarget() []string
	GetRunner() string
	GetTarget() string
	GetVerbose() bool
}
