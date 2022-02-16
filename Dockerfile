FROM golang:1.17.7-alpine3.15

WORKDIR /opt/app/
COPY . .

RUN go mod download && go mod verify

CMD ["go", "run", "main.go"]
