FROM golang:latest

WORKDIR $GOPATH/src/github.com/deozza/golang-blog-native

COPY ./app .

RUN go mod init github.com/deozza/golang-blog-native
RUN go get -u -d -v ./...
RUN go build

CMD ["./golang-blog-native"]