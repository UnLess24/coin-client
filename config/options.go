package config

type options struct {
	configName string
	configPath string
	configType string
}

type Option func(*options)

func WithConfigName(name string) Option {
	return func(o *options) {
		o.configName = name
	}
}

func WithConfigPath(path string) Option {
	return func(o *options) {
		o.configPath = path
	}
}

func WithConfigType(t string) Option {
	return func(o *options) {
		o.configType = t
	}
}
