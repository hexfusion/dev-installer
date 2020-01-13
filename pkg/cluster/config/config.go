package config

type Auth struct {
	Auth  string `yaml:"auth"`
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
}

type Config struct {
	BaseDir `yaml:"base_dir"`
	SSHPath `yaml:"ssh_path"`
	Tokens  []Auth `yaml:"tokens"`
}
