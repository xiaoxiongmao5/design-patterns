package functionaloptions

type Config struct {
	Option1 string
	Option2 int
	// ... other options
}

type Option func(*Config)

func WithOption1(opt string) Option {
	return func(c *Config) {
		c.Option1 = opt
	}
}

func WithOption2(opt int) Option {
	return func(c *Config) {
		c.Option2 = opt
	}
}

func NewConfig(opts ...Option) *Config {
	cfg := &Config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}
