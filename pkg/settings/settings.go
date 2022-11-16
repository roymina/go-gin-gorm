package settings

import (
	"log"

	"gopkg.in/ini.v1"
)

var Cfg *ini.File

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("fatal to parse 'conf/app.ini':%v", err)
	}
}
