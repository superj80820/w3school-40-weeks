#! /bin/bash

DOMAIN=$1
cp /etc/letsencrypt/live/$DOMAIN/privkey.pem privkey.pem
cp /etc/letsencrypt/live/$DOMAIN/cert.pem cert.pem
cp /etc/letsencrypt/live/$DOMAIN/fullchain.pem fullchain.pem
chmod 777 ./privkey.pem
aws iam upload-server-certificate --server-certificate-name $DOMAIN
                                  --certificate-body file://cert.pem
                                  --certificate-chain file://fullchain.pem
                                  --private-key file://privkey.pem