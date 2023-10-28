#!/bin/bash
openssl genrsa -out ./private.key 2048

openssl req \
-new \
-x509 \
-config cert.cnf \
-key private.key \
-out cert.crt \
-days 365 \
-batch


