#!/bin/bash

export BOSH_CLIENT_SECRET=d8hcm8gj6td59as3rliu
export BOSH_CLIENT=admin
bosh -e hackaton -d minecraft1 deploy ./minecraft.yml -v deployment-name=minecraft1 --vars-store ./mincraft1-creds.yml

bosh -e hackaton -d minecraft1 vms

exit 0
