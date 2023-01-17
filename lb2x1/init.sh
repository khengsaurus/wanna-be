#!/bin/bash

# 2 upstream services pointing to the same app
bash scripts/init_apps.sh 1
bash scripts/init_nginx.sh lb2x1