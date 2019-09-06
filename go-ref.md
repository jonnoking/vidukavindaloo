# Go Ref

## Startup

Startup vidukavindaloo-go-www

Connect to network
```docker network connect cloudjonno-network cloudjonno-nginx```

Startup nginx

Connect to network

```docker network connect cloudjonno-network cloudjonno-nginx```

## General

Clean Architecture using Go - https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f

## Go

**Environment Variables**

https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f

$ go get github.com/joho/godotenv

## Go Research

* context
* waitgroups

## Docker


```docker-compose -f nginx-proxy-compose.yaml up -d```

```docker-compose -f viduka-app-compose.yaml up -d --build```

Docker setup for nginx proxy and letsencrypt

https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker-and-nginx-on-ubuntu-18-04

https://github.com/jwilder/nginx-proxy

https://github.com/JrCs/docker-letsencrypt-nginx-proxy-companion

http://oskarhane.com/nginx-as-a-reverse-proxy-in-front-of-your-docker-containers/


### Networking docker containers

```docker network create cloudjonno-network```

```docker network connect cloudjonno-network vidukavindaloo-go-www```

```docker network connect cloudjonno-network cloudjonno-nginx```

```docker exec -ti cloudjonno-nginx ping vidukavindaloo-go-www```

### Logs

```docker logs --tail 50 --follow --timestamps [container name]```

## Bash

How do I reload .bashrc without logging out and back in?
https://stackoverflow.com/questions/2518127/how-do-i-reload-bashrc-without-logging-out-and-back-in



## Links

### Environment variables
https://dev.to/craicoverflow/a-no-nonsense-guide-to-environment-variables-in-go-a2f


### HTTPS

https://blog.kowalczyk.info/article/Jl3G/https-for-free-in-go-with-little-help-of-lets-encrypt.html

https://getgophish.com/blog/post/2018-12-02-building-web-servers-in-go/

https://github.com/jordan-wright/unindexed


https://www.thegeekstuff.com/2017/05/nginx-location-examples/

### Templating

https://meshstudio.io/blog/2017-11-06-serving-html-with-golang/

https://github.com/meshhq/golang-html-template-tutorial/blob/master/main.go


## JSON / Map --> Structure

https://godoc.org/github.com/mitchellh/mapstructure

https://stackoverflow.com/questions/51813411/golang-mapstructure-not-working-as-expected

https://stackoverflow.com/questions/22343083/json-unmarshaling-with-long-numbers-gives-floating-point-number

https://yourbasic.org/golang/json-example/


## Best Practices

https://github.com/golang-standards/project-layout

https://github.com/tmrts/go-patterns




server {
    listen 80;
    server_name cloudjonno.com;
    server_tokens off;

    location /.well-known/acme-challenge/ {
        root /var/www/certbot;
    }

    location / {
        return 301 https://$host$request_uri;
    }
}

server {
    listen 443 ssl;
    server_name cloudjonno.com;
    server_tokens off;

    ssl_certificate /etc/letsencrypt/live/cloudjonno.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/cloudjonno.com/privkey.pem;
    include /etc/letsencrypt/options-ssl-nginx.conf;
    ssl_dhparam /etc/letsencrypt/ssl-dhparams.pem;

    location / {
        proxy_pass  http://vidukavindaloo-go-www:8111;
        proxy_set_header    Host                $http_host;
        proxy_set_header    X-Real-IP           $remote_addr;
        proxy_set_header    X-Forwarded-For     $proxy_add_x_forwarded_for;
    }
}



## Setting up docker images for cloudjonno.com for the first time.

Post: https://medium.com/@pentacent/nginx-and-lets-encrypt-with-docker-in-less-than-5-minutes-b4b8a60d3a71

Github: https://github.com/wmnnd/nginx-certbot

Make changes to `init-letsencrypt.sh` - change example.org to cloudjonno.com

Remove this section for cert validation
```location / {
        proxy_pass  http://vidukavindaloo-go-www:8111;
        proxy_set_header    Host                $http_host;
        proxy_set_header    X-Real-IP           $remote_addr;
        proxy_set_header    X-Forwarded-For     $proxy_add_x_forwarded_for;
    }```

Run `./init-letsencrypt.sh`

pull down docker images

add redirect section to vidukavindaloo image & port

`docker-compose up -d`

connect to network of vidukavindaloo-go-www - make sure vv is alread connected to network

`docker network connect cloudjonno-network cloudjonno-nginx`

