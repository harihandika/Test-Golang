FROM golang:1.18

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o test-golang

CMD [ "go", "run", "app/main.go" ]