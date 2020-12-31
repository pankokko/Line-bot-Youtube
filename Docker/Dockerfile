FROM golang:latest

RUN mkdir /go/src/go-study
WORKDIR /go/src/go-study
#コンテナの中に入った時のルートパス
ENV GO111MODULE=on
COPY . .
RUN go get -u github.com/cosmtrek/air

CMD go run main.go