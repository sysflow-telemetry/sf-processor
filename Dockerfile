FROM golang:1.14

RUN GO111MODULE=on go get -u -v github.com/dgraph-io/dgo/v2

COPY entrypoint.sh .

CMD ./entrypoint.sh
