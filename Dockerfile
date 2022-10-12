FROM golang:alpine

WORKDIR /app
COPY . .

RUN go build -o rest

EXPOSE 8080

CMD [ "./rest" ]