FROM golang:1.18

COPY . .

RUN mkdir /app

WORKDIR /app

COPY . .

CMD [ "harihandika/test-golang" ]