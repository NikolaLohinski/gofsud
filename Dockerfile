ARG GO_IMAGE_VERSION
ARG DISTROLESS_IMAGE
ARG DISTROLESS_VERSION

# Golang
FROM    golang:$GO_IMAGE_VERSION AS essentials
RUN     set -eux; \
        apk update && \
        apk add --no-cache upx git gcc musl-dev
RUN     set -eux; \
        git clone https://github.com/magefile/mage .mage \
        && cd .mage \
        && go run bootstrap.go \
        && cd .. \
        && rm -rf .mage

# Tools
FROM    essentials AS tools
WORKDIR /tmp/src/
COPY    tools tools
RUN     cd tools && mage

# Dependencies
FROM    tools AS dependencies
COPY    go.mod .
RUN     set -eux;\
        go mod download
COPY    spellbook spellbook
COPY    magefile.go magefile.go
COPY    .golangci.yml .golangci.yml

# Source
FROM    dependencies AS source
COPY    . .
RUN     set -eux;\
        go mod vendor
RUN     mage

# Builder
FROM    source AS builder
ARG     VCS_TAG
ENV     VCS_TAG=$VCS_TAG
RUN     mage bin:build
RUN     set -eux; \
        upx -9 .local/bin/gofsud && \
        chmod +x .local/bin/gofsud

# App
FROM       $DISTROLESS_IMAGE:$DISTROLESS_VERSION
COPY       --from=builder /tmp/src/.local/bin/gofsud /app/gofsud
ENTRYPOINT [ "/app/gofsud" ]
