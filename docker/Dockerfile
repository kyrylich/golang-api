FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -v -o server

FROM alpine

COPY --from=builder /app/config/ /app/config/
COPY --from=builder /app/server /app/server

RUN ls /app

EXPOSE 8080

CMD ["/app/server"]