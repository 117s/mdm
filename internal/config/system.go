package config

type System struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`

	// 数据库类型:sqlite|sqlserver|postgresql|mysql
	DbType string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`

	StaticPath     string `mapstructure:"static-path" json:"staticPath" yaml:"static-path"`
	EnableAuditLog bool   `mapstructure:"enable-audit-log" json:"enableAuditLog" yaml:"enable-audit-log"`
}
