#!/bin/bash
## generate CA
# openssl req -x509 -config openssl-ca.cnf -newkey rsa:4096 -sha256 -nodes -out cacert.pem -outform PEM -days 3600
## inspect CA
# openssl x509 -in ca/cacert.pem -text -noout
## Inspect the CA Self Signed Certificate for its Purpose/Ability
# openssl x509 -purpose -in cacert.pem -inform PEM
## Client: Generate Private Key & Certificate Signing Request (CSR)
# openssl req -config client.cnf -newkey rsa:2048 -sha256 -nodes -out clientcert.csr -outform PEM
## Inspect the CSR (Certificate Signing Request)
# openssl req -text -noout -verify -in clientcert.csr
## CA: Sign the CSR
# openssl ca -config tools/ca.cnf -policy signing_policy -extensions signing_req -out clientcert.pem -infiles clientcert.csr

## curl --cert data/certs/clientcert/clientcert.pem --key data/certs/clientcert/clientkey.pem --cacert data/certs/ca/cacert.pem --capath data/certs/ca https://localhost:443

openssl ca -config ca.cnf -policy signing_policy -extensions signing_req -out clientcert.pem -infiles clientcert.csr
openssl req \
-new \
-x509 \
-config cert.cnf \
-key private.key \
-out cert.crt \
-days 365 \
-batch


