package data

var (
	ConfigData *Config = new(Config)
)

type ImapConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Security string `yaml:"security"`
}

type SmtpConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Security string `yaml:"security"`
}

type Config struct {
	Imap []ImapConfig `yaml:"imap"`
	Smtp []SmtpConfig `yaml:"smtp"`
}
