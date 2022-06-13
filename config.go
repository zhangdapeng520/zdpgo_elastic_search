package zdpgo_elasticsearch

type Config struct {
	Debug          bool     `yaml:"debug" json:"debug"`                 // 是否为Debug模式
	LogFilePath    string   `yaml:"log_file_path" json:"log_file_path"` // 日志路径
	Addresses      []string `yaml:"addresses" json:"addresses"`
	Username       string   `yaml:"username" json:"username"`               // 用户名
	Password       string   `yaml:"password" json:"password"`               // 密码
	CertPath       string   `yaml:"cert_path" json:"cert_path"`             // 密码
	IsOpenSniff    bool     `yaml:"is_open_sniff" json:"is_open_sniff"`     // 是否开启sniff模式
	IsCloseZip     bool     `yaml:"is_close_zip" json:"is_close_zip"`       // 是否关闭ZIP压缩
	HealthInterval int      `yaml:"health_interval" json:"health_interval"` // 健康检查时间
	RetryTimeNum   int      `yaml:"retry_time_num" json:"retry_time_num"`   // 失败重试次数
}
