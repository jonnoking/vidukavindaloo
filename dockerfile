FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/vv
COPY . .
COPY .env .
#RUN ls * -la
#RUN pwd
#RUN echo $GOPATH
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./vv-app ./main.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
# make sure working dir is correct to ensure static files are found
WORKDIR /go/bin/vv
COPY --from=build /go/src/vv /go/bin/vv
EXPOSE 8111
ENTRYPOINT /go/bin/vv/vv-app