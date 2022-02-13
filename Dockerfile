FROM golang:1.17.7-alpine3.15

WORKDIR /opt/app
COPY . .

RUN go mod download && go mod verify

EXPOSE 3000

CMD ["go", "run ."]
