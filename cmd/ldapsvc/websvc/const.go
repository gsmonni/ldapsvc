package websvc

import "fmt"

const (
	URIPrefix = "/api/v1/ldap"

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
	URIHealth    = fmt.Sprintf("/ldap/health")
	URIStop      = fmt.Sprintf("/ldap/stop")
)
