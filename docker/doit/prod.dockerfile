FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/Tinee/doit

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo ./cmd/doit_srv/

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/src/github.com/Tinee/doit/doit_srv .

CMD ./doit_srv