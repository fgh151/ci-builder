Requirements
------------

Docker

#Installation

##DNS server
```shell
docker run --detach --restart=always --name dns \
-p 53:53/tcp \
-p 53:53/udp \
--cap-add=NET_ADMIN \
andyshinn/dnsmasq:2.76 --keep-in-foreground --address=/......box/127.0.0.1 --server=1.1.1.1
```
##Nginx
```shell
docker run --detach --restart=always --name nginx \
-v ./nginx.conf:/etc/nginx/nginx.conf:ro \
-p 80:80 \
nginx
```
##Router
```shell
docker network create router

docker run --detach --restart=always --name router --network=router \
-v /var/run/docker.sock:/var/run/docker.sock \
-p 80:8090 \
-l traefik.enable=true \
-l traefik.frontend.rule=Host:router.box \
-l traefik.port=8080 \
traefik:v1.7.15-alpine --api --docker --docker.exposedbydefault=false
```
