FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR ~/go/src/github.com/jonnoking/vidukavindaloo
COPY . .
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build ~/go/src/github.com/jonnoking/vidukavindaloo /go/bin
EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80