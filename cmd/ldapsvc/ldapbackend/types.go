package ldapbackend

import "github.com/go-ldap/ldap"

type (
	LDAPParameters struct {
		BindString   string `json:"bind-string"`
		BindPassword string `json:"bind-password"`
		BaseDN       string `json:"base-dn"`
		Server       string `json:"server"`
		Port         uint16 `json:"port"`
		UseLDAPS     bool   `json:"use-ldaps"`
		Mock         bool   `json:"mock"`
		MockDataFile string `json:"mock-data-file"`
		MockItemsNum uint16 `json:"mock-items-num"`
	}
	QueryResult struct {
		CN        string   `json:"cn,omitempty"`
		UID       string   `json:"uid,omitempty"`
		ClientId  string   `json:"client-id,omitempty"`
		Country   string   `json:"country,omitempty"`
		Groups    []string `json:"groups,omitempty"`
		Roles     []string `json:"roles,omitempty"`
		FirstName string   `json:"first-name,omitempty"`
		LastName  string   `json:"last-name,omitempty"`
	}
	Results []QueryResult

	IProvider interface {
		Connect() error
		Query(query string) (Results, error)
	}
	Provider struct {
		r *Results
		c *ldap.Conn

		parameters *LDAPParameters
	}
)
