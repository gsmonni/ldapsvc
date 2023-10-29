package ldapbackend

import (
	"fmt"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"log"
)

func New(par LDAPParameters) (*Provider, error) {
	p := new(Provider)
	p.r = new(Results)

	p.parameters = par
	if p.parameters.Mock {
		if !common.FileExists(p.parameters.MockDataFile) {
			p.r = GenerateMockData(p.parameters.MockItemsNum)
			if err := SaveResult(p.r, p.parameters.MockDataFile); err != nil {
				return nil, err
			}
		} else {
			if err := common.ReadJson(p.parameters.MockDataFile, p.r); err != nil {
				return nil, fmt.Errorf("cannot load LDAP mock-data from file (%v)", err.Error())
			}
			log.Printf("loaded %d items from  mock-data-file %s", len(*p.r), par.MockDataFile)
		}
		return p, nil
	} else {
		return nil, fmt.Errorf("accessing non-mocked ldap-backend is currently NOT supported")
	}
}

func (p *Provider) Query(q string) (*Results, error) {
	if p.parameters.Mock {
		return QueryMockData(q, p.r)
	} else {
		return nil, fmt.Errorf("accessing non-mocked ldap-backend is currently NOT supported")
	}
}
