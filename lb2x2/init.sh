#!/bin/bash

# Start 2x nginx balancing 2x app each
bash scripts/init_apps.sh 4
bash scripts/init_nginx.sh lb2x2