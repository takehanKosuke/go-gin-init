FROM golang:1.14

COPY . go/src/app_name

WORKDIR go/src/app_name

RUN GOOS=linux GOARCH=amd64 go build main.go

CMD ["./main"]
