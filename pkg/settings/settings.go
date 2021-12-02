package settings

import (
	"github.com/go-ini/ini"
	"github.com/solo-kingdom/meta/pkg/constants"
	"github.com/solo-kingdom/meta/pkg/utils"
	"log"
)

type Server struct {
	Port int
}
type App struct {
	UploadPath string
}

var cfg *ini.File

var ServerConfig = &Server{}
var AppConfig = &App{}
var SectionConfig = map[string]interface{}{
	"server": ServerConfig,
	"app":    AppConfig,
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

	err = utils.EnsureDir(AppConfig.UploadPath)
	if err != nil {
		log.Fatalf("set up upload path failed. [UploadPath=%v]", AppConfig.UploadPath)
	}
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("map conf failed. [section=%s, err=%v]", section, err)
	}
}
