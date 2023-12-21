package cli

import (
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/websvc"
	"log"
	"os"
	"path/filepath"
)

var (
	ConfFile    = filepath.Join("data", "conf", "conf.json")
	DefConfFile = filepath.Join("data", "conf", "default.json")
)

func Parse(cfg *websvc.Parameters) error {
	common.Datapath = os.Getenv("DATAPATH")

	cfg.Build = "1.0.0"
	cfg.Desc = "LDAP Service"
	myConf := filepath.Join(common.Datapath, ConfFile)

	if err := common.ReadJson(myConf, cfg); err != nil {
		log.Printf("error reading configuration file %s (%v)", myConf, err.Error())
		myConf = filepath.Join(common.Datapath, DefConfFile)
		if err := common.ReadJson(myConf, cfg); err != nil {
			log.Printf("error reading configuration file %s (%v)", myConf, err.Error())
		} else {
			log.Printf("configuration loaded from %s", myConf)
		}
	} else {
		log.Printf("configuration loaded from %s", myConf)
	}

	if s, err := conf.Parse("ldapsvc", cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(s)
		}
		return err
	}
	cfg.LDAP.MockDataFile = filepath.Join(common.Datapath, cfg.LDAP.MockDataFile)
	fmt.Println(conf.String(cfg))

	if cfg.SaveLastConfig {
		if err := common.SaveJson(ConfFile, *cfg); err != nil {
			log.Printf("error saving conf %v", err.Error())
		}
	}
	return nil
}
