package optionalinterface

type Config struct {
	Option1 string
	Option2 int
	// ... other options
}

type Optioner interface {
	Apply(*Config)
}

type Option1 struct {
	Value string
}

func (o Option1) Apply(c *Config) {
	c.Option1 = o.Value
}

type Option2 struct {
	Value int
}

func (o Option2) Apply(c *Config) {
	c.Option2 = o.Value
}

func NewConfig(opts ...Optioner) *Config {
	cfg := &Config{}
	for _, opt := range opts {
		opt.Apply(cfg)
	}
	return cfg
}
