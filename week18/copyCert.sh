#! /bin/bash

DOMAIN=$1
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem privkey.pem
cp /etc/letsencrypt/live/$DOMAIN/cert.pem cert.pem
chmod 777 ./privkey.pem