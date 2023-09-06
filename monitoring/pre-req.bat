echo "install buf"
curl -sSL "https://github.com/bufbuild/buf/releases/download/v1.26.1/buf-Windows-x86_64.exe" -o buf.exe
mv buf.exe %GOBIN%\buf.exe
