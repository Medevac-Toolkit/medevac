FROM golang:latest as builder

WORKDIR /app
COPY . .
RUN go build -o medevac main.go

FROM gcr.io/distroless/base
COPY --from=builder /app/medevac /usr/local/bin/medevac
CMD ["medevac"]
