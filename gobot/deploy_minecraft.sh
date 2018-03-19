#!/bin/bash

export BOSH_CLIENT_SECRET=d8hcm8gj6td59as3rliu
export BOSH_CLIENT=admin
bosh -e hackaton -d minecraft1 deploy /home/ubuntu/hackaton/minecraft.yml -v deployment-name=minecraft1 --vars-store /home/ubuntu/hackaton/minecraft1-creds.yml -n

bosh -e hackaton -d minecraft1 vms

exit 0
