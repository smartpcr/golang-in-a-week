FROM alpine:3.18.0

ARG TARGETARCH

RUN apk add --no-cache \
    ca-certificates \
    iptables \
    && update-ca-certificates
RUN apk add --update bash curl

WORKDIR /app

# Create .config directory, used by Delve to store settings, breakpoints, and other debugging-related information.
RUN mkdir .config && \
    chown -R 65532:65532 .config && \
    chmod 700 .config

COPY config.yaml /app
COPY --chown=65532:65532 bin/server /app
COPY --chown=65532:65532 certs/rootCA.crt /app/certs/rootCA.crt
COPY --chown=65532:65532 certs/server.crt /app/certs/server.crt
COPY --chown=65532:65532 certs/server.key /app/certs/server.key
COPY bin/dlv /app
USER 65532:65532

# expose dlv at 2160
EXPOSE 2160

CMD ["/app/dlv", "--listen=:2160", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/server"]
