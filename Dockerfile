FROM golang:1.19
RUN mkdir -p /usr/app
WORKDIR /usr/app
COPY . /usr/app
ENV GOPROXY="https://goproxy.io"
ENV GIN_MODE=release
RUN make build
EXPOSE 8080
ENTRYPOINT ./bin/app 