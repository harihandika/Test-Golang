FROM golang:1.18

COPY . .

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o test-golang

CMD [ "./test-golang" ]