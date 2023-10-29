package ldapbackend

type (
	LDAPParameters struct {
		BindString   string `conf:"default:cn=admin,dc=example,dc=example" json:"bind-string"`
		BindPassword string `conf:"default:mypass,mask" json:"bind-password"`
		Server       string `conf:"default:localhost" json:"server"`
		Port         int    `conf:"default:636" json:"port"`
		UseLDAPS     string `conf:"default:false" json:"use-ldaps"`
	}
)
