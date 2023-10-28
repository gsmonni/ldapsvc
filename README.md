# ldaspvc
this is a simple microservice that pulls LDAP informations and presents them via a URI

- _/api/v1/ldap/health_: returns the service health status (web frontend and LDAP backend)
- /_api/v1/ldap/query/{query-attribute-type}/{query-attribute-value}_: performs a LDAP query to select all objects whose type corresponds to the given value
- _/api/v1/ldap/stop_: stop the webservice

## To run the service
- make run
- _curl http://localhost:8080/api/v1/ldap/health_

## To build the service
- make build
