# Use build arguments to specify the architecture
ARG TARGETARCH

FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Use the build argument for the target architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -a -installsuffix cgo -o main .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./main"]