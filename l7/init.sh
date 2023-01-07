#!bin/bash

# Start 1x nginx and 3x app
bash scripts/init_apps.sh 3
bash scripts/init_nginx.sh l7
