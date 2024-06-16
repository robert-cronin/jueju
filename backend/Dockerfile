# Use build arguments to specify the architecture
ARG TARGETARCH

FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Use the build argument for the target architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -a -installsuffix cgo -o main .

# development
FROM golang:latest AS dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

EXPOSE 3000

ENTRYPOINT ["main"]

# production

FROM alpine:latest  

WORKDIR /app/

COPY --from=builder /app/main .

ENV ENV=production

EXPOSE 3000

CMD ["./main"]