FROM golang:1.16-alpine as build-env

ENV GIN_MODE=release

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN GIN_MODE=release CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /diglett

FROM scratch

WORKDIR /

COPY --from=build-env /diglett /diglett

EXPOSE 8080

CMD ["/diglett"]

