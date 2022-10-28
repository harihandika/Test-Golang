FROM golang:1.18

COPY . .

RUN mkdir /app

WORKDIR /app

COPY . .

CMD [ "./test-golang" ]