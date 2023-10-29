package ldapbackend

import (
	"fmt"
	"github.com/go-ldap/ldap"
	"github.com/gsmonni/ladapsvc/cmd/ldapsvc/common"
	"log"
)

func New(par LDAPParameters) (*Provider, error) {
	p := new(Provider)
	p.r = new(Results)
	p.parameters = par

	if p.parameters.Mock {
		if err := p.mockProvider(); err != nil {
			return nil, fmt.Errorf("error while creating mock-data (%v)", err)
		}
		return p, nil
	} else {
		return nil, fmt.Errorf("accessing non-mocked ldap-backend is currently NOT supported")
	}
}

func (p *Provider) Connect() error {
	if p.parameters.UseLDAPS {
		return fmt.Errorf("LDAPS currently not supported")
	}
	proto := "ldap"
	ldapURL := fmt.Sprintf("%s://%s:%s", proto, p.parameters.Server, p.parameters.Port)
	var err error
	if p.c, err = ldap.DialURL(ldapURL); err != nil {
		return err
	}
	defer p.c.Close()
	return p.bind()
}

func (p *Provider) bind() error {
	if err := p.c.Bind(p.parameters.BindString, p.parameters.BindPassword); err != nil {
		return err
	}
	return nil
}

func (p *Provider) mockProvider() error {
	if !common.FileExists(p.parameters.MockDataFile) {
		p.r = GenerateMockData(p.parameters.MockItemsNum)
		if err := SaveResult(p.r, p.parameters.MockDataFile); err != nil {
			return err
		}
	} else {
		if err := common.ReadJson(p.parameters.MockDataFile, p.r); err != nil {
			return fmt.Errorf("cannot load LDAP mock-data from file (%v)", err.Error())
		}
		log.Printf("loaded %d items from  mock-data-file %s", len(*(p.r)), p.parameters.MockDataFile)
	}
	return nil
}

func (p *Provider) Query(q string) (*Results, error) {
	if p.parameters.Mock {
		return QueryMockData(q, p.r)
	}
	// connect code comes here
	baseDN := p.parameters.BaseDN
	filter := fmt.Sprintf("(%s)", ldap.EscapeFilter(q))

	// Filters must start and finish with ()!
	searchReq := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, 0, 0, 0, false,
		filter,
		[]string{"sAMAccountName"}, []ldap.Control{})

	if result, err := p.c.Search(searchReq); err != nil {
		return nil, fmt.Errorf("failed to query LDAP: %w", err)
	} else {
		return convertResult(result)
	}

}

func convertResult(r *ldap.SearchResult) (*Results, error) {
	return nil, nil
}
