#! /bin/bash

DOMAIN=$1
FOLDER=$2
ARN=$3
aws acm import-certificate --certificate file://$FOLDER/cert.pem --certificate-chain file://$FOLDER/fullchain.pem --private-key file://$FOLDER/privkey.pem --certificate-arn $ARN