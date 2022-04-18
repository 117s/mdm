package config

type Sqlite struct {
	Path   string `mapstructure:"path" json:"path" yaml:"path"`
	Memory bool   `mapstructure:"memory" json:"memory" yaml:"memory"`
}

func (p *Sqlite) Dsn() string {
	if p.Memory == true {
		return ":memory:?cache=shared"
	}
	return p.Path
}
