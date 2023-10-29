# ldaspvc
![image](https://github.com/gsmonni/ldapsvc/assets/142036606/187f6d78-eb55-4aab-9230-402d7f98c904)

## To build the service
- make build (ldapsvc binary is generated inside release/)

## To run the service (without building)
1. make run (this will start the service listening on all interfaces on port 8080)
2. use swagger: _curl http://localhost:8080/swaggerui/_

## CLI parameters
ldapsvc parameters can be set either via environment variables, command-line parameters, or via JSON configuration file located in [_data/conf/conf.json_](data/conf/conf.json)

```azure
go run cmd/main.go -h
Usage: main [options] [arguments]

OPTIONS
  --certificate-use-tls/$LDAPSVC_CERTIFICATE_USE_TLS      <bool>    (default: false)
  --certificate-use-mtls/$LDAPSVC_CERTIFICATE_USE_MTLS    <bool>    (default: false)
  --certificate-strict/$LDAPSVC_CERTIFICATE_STRICT        <bool>    (default: false)
  --certificate-ca-file/$LDAPSVC_CERTIFICATE_CA_FILE      <string>  (default: ca.cert)
  --certificate-cert-file/$LDAPSVC_CERTIFICATE_CERT_FILE  <string>  (default: /Users/gianstefanomonni/git/ladapsvc/data/certs/cert.crt)
  --certificate-key-file/$LDAPSVC_CERTIFICATE_KEY_FILE    <string>  (default: /Users/gianstefanomonni/git/ladapsvc/data/certs/private.key)
  --certificate-ca-path/$LDAPSVC_CERTIFICATE_CA_PATH      <string>  (default: .)
  --local-address/$LDAPSVC_LOCAL_ADDRESS                  <string>  
  --port/$LDAPSVC_PORT                                    <int>     (default: 8080)
  --help/-h display this help message
```
