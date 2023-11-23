FROM nginx:stable-alpine
 
# any steps after this, will execute as if within this directory
WORKDIR /etc/nginx/conf.d
 
COPY docker/dev/nginx/nginx_server.conf ./default.conf
