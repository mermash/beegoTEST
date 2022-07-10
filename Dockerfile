FROM golang:1.18.2-alpine3.16

RUN go install github.com/beego/bee/v2@latest

RUN mkdir -p home/user

COPY . /home/user

WORKDIR /home/user

CMD $GOPATH/bin/bee run -gendoc=true -downdoc=true