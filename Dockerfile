FROM golang:1.14
WORKDIR /go/src/github.com/zillani/qudash
COPY . .
RUN go mod download
EXPOSE 8080
RUN GO111MODULE=off go get github.com/beego/bee
ENTRYPOINT bee run --gendoc=true --downdoc=true

