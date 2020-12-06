FROM golang:1.14

COPY ./api go/src/app_name/api
COPY ./go.mod go/src/app_name/

WORKDIR go/src/app_name

RUN GOOS=linux GOARCH=amd64 go build ./api/main.go ./api/wire_gen.go

CMD ["./main"]
