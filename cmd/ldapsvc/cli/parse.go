package cli

import (
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/websvc"
)

var versionWanted = errors.New("version wanted")

func Parse(cfg *websvc.Parameters) error {
	cfg.Version = conf.Version{
		Build: "1.0.0",
		Desc:  "LDAP Service",
	}

	if s, err := conf.Parse("ldapsvc", cfg); err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(s)
		}
		if errors.Is(err, versionWanted) {
			fmt.Println("%v", cfg.Version)
		}

		return err
	}

	fmt.Println(conf.String(cfg))
	return nil
}
