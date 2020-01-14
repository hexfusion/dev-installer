package config

type Auth struct {
	Auth  string `yaml:"auth"`
	Email string `yaml:"email"`
	Name  string `yaml:"name"`
	Registry string `yaml:"registry"`
}

type File struct {
	BaseDir string `yaml:"basedir"`
	SSHPath string `yaml:"sshpath"`
	Tokens  []Auth `yaml:"tokens"`
}
