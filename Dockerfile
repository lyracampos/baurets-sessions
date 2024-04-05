FROM golang:1.22.1 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /baurets-sessions

FROM alpine

WORKDIR /

COPY --from=build /baurets-sessions /baurets-sessions

EXPOSE 8080

ENTRYPOINT [ "/baurets-sessions" ]
