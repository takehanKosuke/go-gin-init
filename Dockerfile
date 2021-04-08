FROM golang:1.14

COPY ./app go/src/app_name/app
COPY ./go.mod go/src/app_name/

WORKDIR go/src/app_name

RUN GOOS=linux GOARCH=amd64 go build ./app/main.go ./app/wire_gen.go

CMD ["./main"]
