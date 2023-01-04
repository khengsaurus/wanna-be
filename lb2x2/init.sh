#!/bin/bash

# Start 2x nginx balancing 2x app each
bash /Users/kheng/git/wanna-be/scripts/init_apps.sh 4
bash /Users/kheng/git/wanna-be/scripts/init_nginx.sh /Users/kheng/git/wanna-be/lb2x2/nginx.conf 8080-8081:8080-8081