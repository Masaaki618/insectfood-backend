FROM golang:1.24-bullseye AS deploy-builder

ENV GOTOOLCHAIN=auto

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app ./cmd/server

# ---------------------------------------------

FROM debian:bullseye-slim AS deploy

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ---------------------------------------------

FROM golang:1.24 AS dev

ENV GOTOOLCHAIN=auto

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

CMD ["air", "-c", ".air.toml"]
