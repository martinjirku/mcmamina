FROM golang:1.23.0 AS watch-base
    # Install Node.js
    RUN apt-get update && apt-get install -y curl
    RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
    RUN apt-get install -y nodejs
    ENV PNPM_HOME="/pnpm"
    ENV PATH="$PNPM_HOME:$PATH"
    RUN corepack enable && corepack use pnpm@10.3.0
    RUN go install github.com/cosmtrek/air@latest
    RUN go install github.com/go-task/task/v3/cmd/task@latest

FROM watch-base AS watch
    COPY . /app
    WORKDIR /app
    RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
    RUN go mod download

FROM node:22.13.1-slim AS node-base
    ENV PNPM_HOME="/pnpm"
    ENV PATH="$PNPM_HOME:$PATH"
    RUN npm install -g corepack && corepack enable && corepack prepare pnpm@10.3.0 --activate
    COPY . /app
    WORKDIR /app

    RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
    RUN pnpm build

FROM golang:1.23.0 AS be-builder
    COPY . /app
    COPY --from=node-base /app/dist /app/dist
    WORKDIR /app

    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcmamina .

FROM scratch
    COPY --from=be-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
    COPY --from=be-builder /app/mcmamina /mcmamina
    EXPOSE 8080
    CMD ["/mcmamina"]
