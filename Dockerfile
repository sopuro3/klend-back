FROM golang:1.21 AS builder

WORKDIR /go/src/app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static
COPY --from=builder /go/bin/app /go/src/app/.env /
EXPOSE 8080
USER nonroot:nonroot
CMD ["/app"]
