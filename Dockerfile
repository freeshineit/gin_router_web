# syntax=docker/dockerfile:1

FROM golang:1.19

WORKDIR /

COPY go.mod ./
COPY go.sum ./
COPY controllers ./controllers
COPY public ./public
COPY router ./router
COPY serialize ./serialize
COPY main.go ./
COPY views ./views

RUN go build -o gin_router_web

EXPOSE 8080

CMD [ "/gin_router_web"]
