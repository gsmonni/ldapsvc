# ldaspvc
LDAPSvc is a golang REST API server that exposes an endpoint to browse a LDAP directory 

![image](https://github.com/gsmonni/ldapsvc/assets/142036606/187f6d78-eb55-4aab-9230-402d7f98c904)

## To build the service
- make build (ldapsvc binary is generated inside release/)

## To run the service (without building)
1. make run (this will start the service listening on all interfaces on port 8080)
2. use swagger: _curl http://localhost:8080/swaggerui/_

### LDAP Data
curently LDAP connection is mocked. Mock data is located within 
[_data/ldapsvc/ldap-data.json_ ](data/ldapsvc/ldap-data.json)
Valid client-id are
- 64975760-cdd6-4910-8133-928ea48bd091
- 256731ab-8c3f-4006-91ad-aa1784d59d0b
![image](https://github.com/gsmonni/ldapsvc/assets/142036606/79c61717-d50e-4a0f-aaed-49e43d653393)

## CLI parameters
ldapsvc parameters can be set either via environment variables, command-line parameters, or via JSON configuration file located in [_data/conf/conf.json_](data/conf/conf.json)

```azure
go run cmd/main.go -h
Usage: main [options] [arguments]

OPTIONS
  --certificate-use-tls/$LDAPSVC_CERTIFICATE_USE_TLS      <bool>    (default: false)
  --certificate-use-mtls/$LDAPSVC_CERTIFICATE_USE_MTLS    <bool>    (default: false)
  --certificate-ca-file/$LDAPSVC_CERTIFICATE_CA_FILE      <string>  (default: ca.cert)
  --certificate-cert-file/$LDAPSVC_CERTIFICATE_CERT_FILE  <string>  (default: data/certs/cert.crt)
  --certificate-key-file/$LDAPSVC_CERTIFICATE_KEY_FILE    <string>  (default: data/certs/private.key)
  --certificate-ca-path/$LDAPSVC_CERTIFICATE_CA_PATH      <string>  (default: .)
  --local-address/$LDAPSVC_LOCAL_ADDRESS                  <string>  
  --port/$LDAPSVC_PORT                                    <int>     (default: 8080)
  --ldap-bind-string/$LDAPSVC_LDAP_BIND_STRING            <string>  (default: cn=admin)
  --ldap-bind-password/$LDAPSVC_LDAP_BIND_PASSWORD        <string>  (default: xxxxxx)
  --ldap-server/$LDAPSVC_LDAP_SERVER                      <string>  (default: localhost)
  --ldap-port/$LDAPSVC_LDAP_PORT                          <int>     (default: 636)
  --ldap-use-ldaps/$LDAPSVC_LDAP_USE_LDAPS                <string>  (default: false)
  --help/-h                                               
  display this help message


```
