@echo off
SETLOCAL ENABLEDELAYEDEXPANSION

:: Check if certs/rootCA.key and certs/rootCA.crt exist
IF NOT EXIST certs\rootCA.key (
  IF NOT EXIST certs\rootCA.crt (
    openssl ^
      req ^
      -new ^
      -newkey rsa:4096 ^
      -days 1024 ^
      -nodes ^
      -x509 ^
      -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" ^
      -keyout certs\rootCA.key ^
      -out certs\rootCA.crt
  )
)

:: Check if certs/server.key and certs/server.crt exist
IF NOT EXIST certs\server.key (
  IF NOT EXIST certs\server.crt (
    openssl ^
      req ^
      -new ^
      -newkey rsa:2048 ^
      -days 372 ^
      -nodes ^
      -x509 ^
      -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" ^
      -addext "subjectAltName=DNS:example.com,DNS:example.net,DNS:otel_collector,DNS:localhost" ^
      -CA certs\rootCA.crt ^
      -CAkey certs\rootCA.key  ^
      -keyout certs\server.key ^
      -out certs\server.crt
  )
)

:: Check if certs/client.key and certs/client.crt exist
IF NOT EXIST certs\client.key (
  IF NOT EXIST certs\client.crt (
    openssl ^
      req ^
      -new ^
      -newkey rsa:2048 ^
      -days 372 ^
      -nodes ^
      -x509 ^
      -subj "/C=US/ST=CA/O=MyOrg/CN=myOrgCA" ^
      -addext "subjectAltName=DNS:example.com,DNS:example.net,DNS:otel_collector,DNS:localhost" ^
      -CA certs\rootCA.crt ^
      -CAkey certs\rootCA.key  ^
      -keyout certs\client.key ^
      -out certs\client.crt
  )
)

:: Remove .csr and .srl files
del /Q certs\*.csr
del /Q certs\*.srl

:: Change file permissions (this is tricky in Windows and may require external tools)
:: For simplicity, this step is commented out
:: Use icacls or other utilities if needed to set permissions

:: icacls certs\rootCA.crt /grant Everyone:R
:: icacls certs\server.crt /grant Everyone:R
:: icacls certs\client.crt /grant Everyone:R

ENDLOCAL
