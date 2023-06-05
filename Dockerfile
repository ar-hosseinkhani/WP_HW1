FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY pkg ./pkg
COPY auth ./auth
COPY biz ./biz
COPY gateway ./gateway

RUN mkdir out

RUN go build -o ./out ./...

EXPOSE 5052

CMD [ "./out/auth" ]
