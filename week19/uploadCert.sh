#! /bin/bash

aws acm import-certificate --certificate file://cert.pem --certificate-chain file://fullchain.pem --private-key file://privkey.pem