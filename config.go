package main

import (
	"github.com/jinzhu/configor"
)

// App config struct
type Config struct {
	Mainnet struct {
		URL      string `default:"" json:"mainneturl" form:"mainneturl" query:"mainneturl"`
		User     string `default:"" json:"mainnetuser" form:"mainnetuser" query:"mainnetuser"`
		Password string `default:"" json:"mainnetpassword" form:"mainnetpassword" query:"mainnetpassword"`
		ChainID  string `default:"" json:"mainnetchainid" form:"mainnetchainid" query:"mainnetchainid"`
	}
	Customnet struct {
		URL      string `default:"" json:"customneturl" form:"customneturl" query:"customneturl"`
		User     string `default:"" json:"customnetuser" form:"customnetuser" query:"customnetuser"`
		Password string `default:"" json:"customnetpassword" form:"customnetpassword" query:"customnetpassword"`
		ChainID  string `default:"" json:"customnetchainid" form:"customnetchainid" query:"customnetchainid"`
	}
}

// Create config from configFile
func NewConfig(configFile string) (*Config, error) {

	config := new(Config)

	if err := configor.Load(config, configFile); err != nil {
		return nil, err
	}
	return config, nil
}
