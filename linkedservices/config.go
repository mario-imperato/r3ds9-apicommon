package linkedservices

import (
	"github.com/mario-imperato/r3ds9-apicommon/linkedservices/kafka"
	"github.com/mario-imperato/r3ds9-apicommon/linkedservices/mongodb"
	"github.com/mario-imperato/r3ds9-apicommon/linkedservices/restclient"
)

type Config struct {
	RestClient *restclient.Config    `mapstructure:"rest-client" json:"rest-client" yaml:"rest-client"`
	Kafka      []kafka.Config        `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
	Mongo      []mongodb.MongoConfig `mapstructure:"mongo" json:"mongo" yaml:"mongo"`
}

func (c *Config) PostProcess() error {

	var err error
	for i, _ := range c.Kafka {
		err = c.Kafka[i].PostProcess()
	}
	if err != nil {
		return err
	}

	return nil
}
