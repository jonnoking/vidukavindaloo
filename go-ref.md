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

