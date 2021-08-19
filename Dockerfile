FROM golang:1.16-alpine as build-env

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /diglett

FROM gcr.io/distroless/base
COPY --from=build-env /diglett /

EXPOSE 8080

CMD [ "/diglett" ]
