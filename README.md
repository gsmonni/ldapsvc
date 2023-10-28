# ldaspvc
this is a simple (http/https) microservice that pulls LDAP informations and presents them via a URI


- _/api/v1/ldap/health_: returns the service health status (web frontend and LDAP backend)
- /_api/v1/ldap/query/{query-attribute-type}/{query-attribute-value}_: performs a LDAP query to select all objects whose type corresponds to the given value
- _/api/v1/ldap/stop_: stop the webservice

## To build the service
- make build (ldapsvc binary is generated inside release/)

## To run the service (without building)
- make run
- _curl http://localhost:8080/api/v1/ldap/health_

## parameters
- server: 
  - address, port
  - TLS, mTLS 
  - certificate-file, key-file, ca, ca-path
- LDAP connection
  - server address, port
  - username, password