Following Hussein Nasser's [Introduction to NGINX](https://www.udemy.com/course/nginx-crash-course/)

```
# nginx with a html path
docker run \
  --name nginx \
  --hostname ng1 \
  -p 80:80 \
  -v html:/usr/share/nginx/html \
  -d \
  nginx

# ssh into a container
docker exec -it <container-name> bash
```