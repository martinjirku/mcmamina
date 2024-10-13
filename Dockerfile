FROM golang:1.23.0 as watch-base
    # Install Node.js
    RUN apt-get update && apt-get install -y curl
    RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
    RUN apt-get install -y nodejs
    ENV PNPM_HOME="/pnpm"
    ENV PATH="$PNPM_HOME:$PATH"
    RUN corepack enable && corepack use pnpm@8.15.6
    RUN go install github.com/air-verse/air@latest
    RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM watch-base as watch
    WORKDIR /app
    COPY . /app
    RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
    RUN pnpm install
    RUN go mod download

FROM node:20-slim AS node-base
    ENV PNPM_HOME="/pnpm"
    ENV PATH="$PNPM_HOME:$PATH"
    RUN corepack enable
    COPY . /app
    WORKDIR /app

    RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
    RUN pnpm build

FROM golang:1.23.0 as be-builder
    COPY . /app
    COPY --from=node-base /app/dist /app/dist
    WORKDIR /app

    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcmamina .

FROM alpine:3.13.5 as migrations
    RUN apk add --no-cache ca-certificates curl
    WORKDIR /app
    ARG TARGETPLATFORM
    RUN ARCH=$(uname -m) && \
        if [ "$ARCH" = "x86_64" ]; then \
            ARCH="amd64"; \
        elif [ "$ARCH" = "aarch64" ]; then \
            ARCH="arm64"; \
        fi && \
        echo ${ARCH} && \
        curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-${ARCH}.tar.gz | tar -xz -C /app -f -
    COPY ./scripts/migrations.sh /app/entrypoint.sh
    RUN chmod +x /app/entrypoint.sh
    COPY ./migrations /migrations
    CMD ["./entrypoint.sh"]

FROM scratch
    COPY --from=be-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
    COPY --from=be-builder /app/mcmamina /mcmamina
    EXPOSE 8080
    CMD ["/mcmamina"]
