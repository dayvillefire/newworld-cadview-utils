package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var (
	Config           *AppConfig
	cachedConfigPath string
)

type CadInstance struct {
	Name     string `yaml:"name"`
	URL      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	FDID     string `yaml:"fdid"`
}

type AppConfig struct {
	Debug         bool                   `yaml:"debug"`
	PollInterval  int                    `yaml:"poll-interval"`
	ReauthMinutes int                    `yaml:"reauth-minutes"`
	PurgeMinutes  int                    `yaml:"purge-minutes"`
	CadInstances  map[string]CadInstance `yaml:"cad"`
	Paths         struct {
		SerializationFile string `yaml:"serial-file"`
	} `yaml:"paths"`
	Discord struct {
		Token     string `yaml:"token"`
		ChannelID string `yaml:"channel-id"`
	} `yaml:"discord"`
	LoggedStatuses []string `yaml:"logged-statuses"`
}

func (c *AppConfig) SetDefaults() {
}

func LoadConfigWithDefaults(configPath string) (*AppConfig, error) {
	cachedConfigPath = configPath
	c := &AppConfig{}
	c.SetDefaults()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal([]byte(data), c)
	return c, err
}
