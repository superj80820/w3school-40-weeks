#! /bin/bash

DOMAIN=$1
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem privkey.pem
cp /etc/letsencrypt/live/$DOMAIN/cert.pem cert.pem
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem fullchain.pem
chmod 777 ./privkey.pem