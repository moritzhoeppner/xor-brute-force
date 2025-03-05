# syntax=docker/dockerfile:1

# https://www.docker.com/blog/containerize-your-go-developer-environment-part-3/

FROM --platform=${BUILDPLATFORM} golang:1.24-alpine AS base
WORKDIR /build
ENV CGO_ENABLED=0

FROM base AS build
ARG TARGETOS
ARG TARGETARCH
RUN --mount=target=. GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/xor-brute-force .

FROM base AS test
RUN --mount=target=. go test ./...

FROM scratch AS bin-unix
COPY --from=build /out/xor-brute-force /

FROM bin-unix AS bin-linux
FROM bin-unix AS bin-darwin

FROM scratch AS bin-windows
COPY --from=build /out/xor-brute-force /xor-brute-force.exe

FROM bin-${TARGETOS} AS bin
