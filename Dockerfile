FROM golang:1.15

WORKDIR $GOPATH/src/github.com/deozza/golang-blog-native

COPY ./app .

#RUN go mod init github.com/deozza/golang-blog-native
#RUN go get -d -v ./...
#RUN go install -v ./...

#EXPOSE 8080

#CMD ["golang-blog-native"]