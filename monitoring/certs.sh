#!/usr/bin/env bash
shopt -s nullglob globstar
set -x # have bash print command been ran
set -e # fail if any command fails

setup_certs(){
  { # create CA.
    if [[ ! -f certs/rootCA.key && ! -f certs/rootCA.crt ]]; then
      openssl \
        req \
        -new \
        -newkey rsa:4096 \
        -days 1024 \
        -nodes \
        -x509 \
        -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" \
        -keyout certs/rootCA.key \
        -out certs/rootCA.crt
    fi
  }

  { # create server certs.
    if [[ ! -f certs/server.key && ! -f certs/server.crt ]]; then
      openssl \
        req \
        -new \
        -newkey rsa:2048 \
        -days 372 \
        -nodes \
        -x509 \
        -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" \
        -addext "subjectAltName=DNS:example.com,DNS:example.net,DNS:otel_collector,DNS:localhost" \
        -CA certs/rootCA.crt \
        -CAkey certs/rootCA.key  \
        -keyout certs/server.key \
        -out certs/server.crt
    fi
  }

  { # create client certs.
    if [[ ! -f certs/client.key && ! -f certs/client.crt ]]; then
      openssl \
        req \
        -new \
        -newkey rsa:2048 \
        -days 372 \
        -nodes \
        -x509 \
        -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" \
        -addext "subjectAltName=DNS:example.com,DNS:example.net,DNS:otel_collector,DNS:localhost" \
        -CA certs/rootCA.crt \
        -CAkey certs/rootCA.key  \
        -keyout certs/client.key \
        -out certs/client.crt
    fi
  }

  { # clean
    rm -rf certs/*.csr
    rm -rf certs/*.srl

    chmod 666 certs/rootCA.crt certs/server.crt certs/client.crt
    chmod 644 certs/rootCA.key certs/server.key certs/client.key
  }
}
setup_certs