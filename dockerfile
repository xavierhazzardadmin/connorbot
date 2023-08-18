FROM golang:1.21.0-alpine3.18 AS BuildStage

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /connorbot

CMD ["/zenbot"]

#Deploy stage

FROM alpine:latest

WORKDIR /

COPY --from=BuildStage /connorbot /connorbot

USER nonroot:nonroot

ENTRYPOINT ["/connorbot"]
