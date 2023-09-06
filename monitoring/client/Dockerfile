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

COPY --chown=65532:65532 bin/client /app
COPY bin/dlv /app
USER 65532:65532

# expose dlv at 2160
EXPOSE 2160

CMD ["/app/dlv", "--listen=:2160", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/app/client"]
