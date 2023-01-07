#!/bin/bash

# Start 1x nginx + 1x app
bash scripts/init_apps.sh 1
bash scripts/init_nginx.sh single