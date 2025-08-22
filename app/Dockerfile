FROM golang:1.25.0-alpine3.22 as builder
WORKDIR /app
COPY . .
RUN go build -o app

FROM scratch
COPY --from=builder /app/app /app
ARG VERSION
ENV VERSION=$VERSION
ENTRYPOINT ["/app"]
