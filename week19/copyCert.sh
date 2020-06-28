#! /bin/bash

DOMAIN=$1
FOLDER=$2
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem $FOLDER/privkey.pem
cp /etc/letsencrypt/live/$DOMAIN/cert.pem $FOLDER/cert.pem
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem $FOLDER/fullchain.pem
chmod 777 $FOLDER/privkey.pem