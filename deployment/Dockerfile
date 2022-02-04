# README run from root dir (./..)
# docker image build -f deployment/DockerFile -t app:1.0 .
# docker container run --publish 8080:8080 --name app app:1.0

FROM golang:1.17 as builder

WORKDIR /go/src/app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN CGO_ENABLED=0 go build -o /go/bin/restapi ./cmd/restapi


FROM alpine:latest

#RUN apk --no-cache add ca-certificates
WORKDIR /bin
COPY config /bin/config
COPY --from=builder /go/bin/restapi .

EXPOSE 81

CMD ["restapi"]
