FROM golang:1.19

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN cd cmd/service; go build -v -o /usr/local/bin/app .

CMD ["app"]