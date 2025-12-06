FROM golang:1.21-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /payment-api ./cmd


FROM scratch
COPY --from=build /payment-api /payment-api
EXPOSE 8080
ENTRYPOINT ["/payment-api"]