package cli

import (
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/websvc"
	"log"
)

const ConfFile = "data/conf/conf.json"
const DefConfFile = "data/conf/default.json"

func Parse(cfg *websvc.Parameters) error {

	cfg.Build = "1.0.0"
	cfg.Desc = "LDAP Service"
	myConf := ""

	if common.FileExists(ConfFile) {
		myConf = ConfFile
	} else {
		myConf = DefConfFile
	}

	if common.FileExists(myConf) {
		log.Printf("loading configuration from %s", ConfFile)
		if err := common.ReadJson(ConfFile, cfg); err != nil {
			log.Printf("error reading configuration file %s", err.Error())
		} else {
			fmt.Println(conf.String(cfg))
		}
	}

	if s, err := conf.Parse("ldapsvc", cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(s)
		}
		return err
	}
	if cfg.SaveLastConfig {
		if err := common.SaveJson(ConfFile, *cfg); err != nil {
			log.Printf("error saving conf %v", err.Error())
		}
	}

	fmt.Println(conf.String(cfg))
	return nil
}
