FROM golang:alpine

WORKDIR /app

COPY . /app/
RUN go mod download

RUN go build -o /rest

EXPOSE 3000

CMD [ "/rest" ]