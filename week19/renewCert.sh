#! /bin/bash

ARN=$1
aws acm import-certificate --certificate file://cert.pem --certificate-chain file://fullchain.pem --private-key file://privkey.pem --certificate-arn $ARN