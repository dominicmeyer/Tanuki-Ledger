FROM golang:1.23.6

RUN apt-get update \
    && apt-get install -y \
        nodejs=18.19.0* \
        npm=9.2.0* \
        # cypress dependencies from https://docs.cypress.io/app/get-started/install-cypress#Linux-Prerequisites
        libgtk2.0-0=2.24.33* \
        libgtk-3-0=3.24.38* \
        libgbm-dev=22.3.6* \
        libnotify-dev=0.8.1* \
        libnss3=2:3.87.1* \
        libxss1=1:1.2.3* \
        libasound2=1.2.8* \
        libxtst6=2:1.2.3* \
        xauth=1:1.1.2* \
        xvfb=2:21.1.7* \
    && go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.5 \
    && go install golang.org/x/tools/cmd/goimports@v0.30.0 \
    && go install github.com/fzipp/gocyclo/cmd/gocyclo@v0.6.0 \
    && go install github.com/air-verse/air@v1.61.7 \
    && go install github.com/a-h/templ/cmd/templ@v0.3.833