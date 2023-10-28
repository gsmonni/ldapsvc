# ldaspvc
this is a simple microservice that pulls LDAP informations and presents them via a URI

- _/api/v1/ldap/health_: returns the service health status (web frontend and LDAP backend)
- /_api/v1/ldap/query/{query-attribute-type}/{query-attribute-value}_: performs a LDAP query to select all objects whose type corresponds to the given value
- _/api/v1/ldap/stop_: stop the webservice

