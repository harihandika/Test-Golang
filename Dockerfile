FROM golang:1.18

RUN mkdir /app

WORKDIR /app

COPY . .

RUN go build -o harihandika/test-golang

CMD [ "./harihandika/test-golang" ]