FROM golang:1.17 AS build

WORKDIR /app

COPY ./ ./

RUN go mod init linha-de-comando

RUN go get golang.org/x/net/html
RUN go get github.com/go-redis/redis
RUN go get github.com/urfave/cli

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]


### UMA FORMA MAIS REDUZIDA
# FROM golang:onbuild

# EXPOSE 8080

### Uma forma mais completa
# FROM golang

# ADD . /go/src/wscraper

# RUN go get golang.org/x/net/html
# RUN go get github.com/go-redis/redis

# RUN go install wscraper

# ENTRYPOINT /go/bin/wscraper

# EXPOSE 8080