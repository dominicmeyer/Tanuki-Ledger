FROM golang:1.23.6

RUN apt-get update \
    && apt-get install -y \
        gcc=4:12.2.0-3 \
        libgl1-mesa-dev=22.3.6-1* \
        xorg-dev=1:7.7+23 \
        libxkbcommon-dev=1.5.0-1 \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5 \
    && go install golang.org/x/tools/cmd/goimports@v0.30.0 \
    && go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.6.0