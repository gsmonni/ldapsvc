package ldapbackend

type (
	LDAPParameters struct {
		BindString   string `conf:"default:cn=admin,dc=example,dc=example" json:"bind-string"`
		BindPassword string `conf:"default:mypass,mask" json:"bind-password"`
		Server       string `conf:"default:localhost" json:"server"`
		Port         uint16 `conf:"default:636" json:"port"`
		UseLDAPS     bool   `conf:"default:false" json:"use-ldaps"`
		Mock         bool   `conf:"default:true" json:"mock"`
		MockDataFile string `conf:"default:data/ldapsvc/ldap-data.json" json:"mock-data-file"`
		MockItemsNum uint16 `conf:"default:20" json:"mock-items-num"`
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

		parameters LDAPParameters
	}
)
