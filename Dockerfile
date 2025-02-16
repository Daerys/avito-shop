FROM golang:1.23.1

WORKDIR ${GOPATH}/avito-shop/
COPY . ${GOPATH}/avito-shop/

RUN go install github.com/easyp-tech/easyp/cmd/easyp@latest

RUN go mod tidy && go mod download && easyp mod update && easyp generate

RUN go build -o /build ./cmd/server \
    && go clean -cache -modcache

EXPOSE 8080

CMD ["/build"]