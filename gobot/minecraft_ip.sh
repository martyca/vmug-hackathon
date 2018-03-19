export BOSH_CLIENT_SECRET=d8hcm8gj6td59as3rliu
export BOSH_CLIENT=admin
bosh -e hackaton -d minecraft1 vms --json | jq .Tables[].Rows[].ips
