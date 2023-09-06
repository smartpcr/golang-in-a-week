#!/usr/bin/env bash
shopt -s nullglob globstar
set -x # have bash print command been ran
set -e # fail if any command fails

setup_certs(){
  { # create CA.
    if [[ ! -f confs/rootCA.key && ! -f confs/rootCA.crt ]]; then
      openssl \
        req \
        -new \
        -newkey rsa:4096 \
        -days 1024 \
        -nodes \
        -x509 \
        -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" \
        -keyout confs/rootCA.key \
        -out confs/rootCA.crt
    fi
  }

  { # create server certs.
    if [[ ! -f confs/server.key && ! -f confs/server.crt ]]; then
      openssl \
        req \
        -new \
        -newkey rsa:2048 \
        -days 372 \
        -nodes \
        -x509 \
        -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" \
        -addext "subjectAltName=DNS:example.com,DNS:example.net,DNS:otel_collector,DNS:localhost" \
        -CA confs/rootCA.crt \
        -CAkey confs/rootCA.key  \
        -keyout confs/server.key \
        -out confs/server.crt
    fi
  }

  { # create client certs.
    if [[ ! -f confs/client.key && ! -f confs/client.crt ]]; then
      openssl \
        req \
        -new \
        -newkey rsa:2048 \
        -days 372 \
        -nodes \
        -x509 \
        -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" \
        -addext "subjectAltName=DNS:example.com,DNS:example.net,DNS:otel_collector,DNS:localhost" \
        -CA confs/rootCA.crt \
        -CAkey confs/rootCA.key  \
        -keyout confs/client.key \
        -out confs/client.crt
    fi
  }

  { # clean
    rm -rf confs/*.csr
    rm -rf confs/*.srl

    chmod 666 confs/rootCA.crt confs/server.crt confs/client.crt
    chmod 644 confs/rootCA.key confs/server.key confs/client.key
  }
}
setup_certs