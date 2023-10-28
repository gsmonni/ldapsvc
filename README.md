# ldaspvc
this is a simple microservice that pulls LDAP informations and presents them via a URI

- /api/v1/ldap/health: returns the healt status
- /api/v1/ldap/query/{query-attribute-type}/{query-attribute-value}: performs a LDAP query to select all objects whose type corresponds to the given value
- /api/v1/ldap/stop: stop the webservice

