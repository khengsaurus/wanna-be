#!bin/bash

# Start 1x nginx and 3x app
bash /Users/kheng/git/wanna-be/scripts/init_apps.sh 3
bash /Users/kheng/git/wanna-be/scripts/init_nginx.sh /Users/kheng/git/wanna-be/l7/nginx.conf 8080:80