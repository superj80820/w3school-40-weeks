#! /bin/bash

DOMAIN=$1
ARN=$2
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem privkey.pem
cp /etc/letsencrypt/live/$DOMAIN/cert.pem cert.pem
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem fullchain.pem
chmod 777 ./privkey.pem

aws acm import-certificate --certificate file://cert.pem --certificate-chain file://fullchain.pem --private-key file://privkey.pem --certificate-arn $ARN