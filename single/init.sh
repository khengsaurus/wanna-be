#!/bin/bash

# Start 1x nginx + 1x app
bash /Users/kheng/git/wanna-be/scripts/init_apps.sh 1
bash /Users/kheng/git/wanna-be/scripts/init_nginx.sh /Users/kheng/git/wanna-be/single/nginx.conf 8080:80