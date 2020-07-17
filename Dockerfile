FROM golang:1.12.9-alpine3.10 as builder
ENV NAME=Sri
COPY main.go .
RUN go build -o /app main.go

FROM alpine:3.10
COPY --from=builder /app .
CMD [ "./app" ]

