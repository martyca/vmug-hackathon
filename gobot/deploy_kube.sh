#!/bin/bash

pks login -a https://uaa.pks.pcf.lab.local -u admin -k -p UjfQ_II0P--mGynVk64ssXyUykA3oGeB
pks create-cluster vmug --external-hostname vmug.pks.pcf.lab.local --plan small

exit 0
