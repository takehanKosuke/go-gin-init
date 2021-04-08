FROM golang:1.14

COPY ./app /app_name/app
COPY ./go.mod /app_name/

WORKDIR /app_name

RUN GOOS=linux GOARCH=amd64 go build ./app/cmd/main.go ./app/cmd/wire_gen.go

CMD ["./main"]
