FROM golang

ADD . /go/src/github.com/aranair/remindbot/

WORKDIR /go/src/github.com/aranair/remindbot
RUN go get ./...
RUN go install

ADD config.toml /go/bin/

ENTRYPOINT /go/bin/remindbot

EXPOSE 8080