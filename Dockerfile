FROM golang:1.23.1 AS build-stage

WORKDIR /src

COPY ./src/ .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GIN_MODE=release go build -o ./release/ ./cmd/api/main.go

# Deploy application binary into a lean image
FROM alpine:3.20.3
#FROM gcr.io/distroless/base-debian11 

WORKDIR /

COPY --from=build-stage /src/release/main /api

EXPOSE 8081

ENTRYPOINT ["/api"]
