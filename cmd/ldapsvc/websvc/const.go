package websvc

import "fmt"

const (
	ServiceStatusUp   = "UP"
	ServiceStatusDown = "DOWN"
	LDAPStatusUp      = "UP"
	LDAPStatusDown    = "DOWN"

	QueryAttributeType  = "query-attribute-type"
	QueryAttributeValue = "query-attribute-value"
)

// these variables should be seen as constants
var (
	URILDAPQuery = fmt.Sprintf("/ldap/query/{%s}/{%s}", QueryAttributeType, QueryAttributeValue)
	URIHealth    = "/ldap/health"
	URIStop      = "/ldap/stop"
)
