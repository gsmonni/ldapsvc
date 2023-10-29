#!/bin/bash
# generate CA
openssl req -x509 -config ca.cnf -newkey rsa:4096 -sha256 -nodes -out ca/cacert.pem -outform PEM -days 3600

openssl genrsa -out ./private.key 2048

openssl req \
-new \
-x509 \
-config cert.cnf \
-key private.key \
-out cert.crt \
-days 365 \
-batch


