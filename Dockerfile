FROM golang:1.25.1-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app ./cmd/server

# ---------------------------------------------

FROM debian:bullseye-slim AS deploy

RUN apt-get update

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ---------------------------------------------

FROM golang:1.25.1 AS dev

WORKDIR /app

RUN go install github.com/air-verse/air@latest
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

CMD ["air", "-c", ".air.toml"]
