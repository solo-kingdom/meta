package settings

import (
	"github.com/go-ini/ini"
	"github.com/solo-kingdom/meta/pkg/constants"
	"log"
)

type Server struct {
	Port int
}

var cfg *ini.File
var ServerConfig = &Server{}
var SectionConfig = map[string]interface{}{
	"server": ServerConfig,
}

func SetUp() {
	var err error
	cfg, err = ini.Load(constants.AppConfig)
	if err != nil {
		log.Fatalf("failed to parse conf. [conf=%v, error=%v]",
			constants.AppConfig, err)
	}

	for section := range SectionConfig {
		mapTo(section, SectionConfig[section])
	}
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("map conf failed. [section=%s, err=%v]", section, err)
	}
}
