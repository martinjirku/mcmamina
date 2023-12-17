FROM node:20-slim AS fe-build

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable
COPY . /app
WORKDIR /app

RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN pnpm build

FROM golang:1.21.5 as be-builder
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . /app
COPY --from=fe-build /app/dist /app/dist
WORKDIR /app

RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mcmamina .


FROM scratch
COPY --from=be-builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=be-builder /app/mcmamina /mcmamina

EXPOSE 3000

CMD ["/mcmamina"]
