package settings

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/go-ini/ini"
	"github.com/solo-kingdom/meta/pkg/constants"
	"github.com/solo-kingdom/meta/pkg/e"
	"github.com/solo-kingdom/meta/pkg/global"
	"github.com/solo-kingdom/meta/pkg/utils"
	"log"
	"os"
	"path"
)

type Server struct {
	Port int
}
type App struct {
	UploadPath string
	KVPath     string
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
		log.Fatalf("failed to parse conf. [conf=%v, e=%v]",
			constants.AppConfig, err)
	}

	for section := range SectionConfig {
		mapTo(section, SectionConfig[section])
	}

	// if set META_HOME
	mh := os.Getenv("META_HOME")
	if len(mh) > 0 {
		AppConfig.UploadPath = path.Join(mh, "upload")
		AppConfig.KVPath = path.Join(mh, "kv")
	}

	err = utils.EnsureDir(AppConfig.UploadPath)
	err = utils.EnsureDir(AppConfig.KVPath)
	if err != nil {
		log.Fatalf("set up upload path failed. [UploadPath=%v]", AppConfig.UploadPath)
	}

	// init kv instance
	global.GV.KV, err = badger.Open(badger.DefaultOptions(AppConfig.KVPath))
	if err != nil {
		log.Fatalf("load kv database failed. [e=%v]", e.ErrMsg(err))
	}
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("map conf failed. [section=%s, err=%v]", section, err)
	}
}
