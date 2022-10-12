FROM golang:alpine

WORKDIR /app
COPY . .

RUN go build -o rest

EXPOSE 8081

CMD [ "./rest" ]