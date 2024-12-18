package duncan

type Appconfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type ConnnectionConfig struct {
	Appconfig
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

type Conections struct {
	Redis    ConnnectionConfig `yaml:"redis"`
	Database ConnnectionConfig `yaml:"database"`
}

type DuncanConfig struct {
	App         Appconfig  `yaml:"app"`
	Connections Conections `yaml:"connections"`
}
