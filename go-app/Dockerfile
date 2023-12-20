FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download
ADD . .

RUN go get github.com/caoquy2000/meeting-app/go-app

EXPOSE 8000

ENTRYPOINT go-app --build="go build main.go" --command=./main
