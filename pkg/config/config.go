package config

type Config interface {
	GetDebounce() string
	GetIgnoreTarget() []string
	GetRunner() string
	GetTarget() string
	GetVerbose() bool
}
