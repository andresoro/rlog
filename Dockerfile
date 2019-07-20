ARG GO_VERSION=1.11

FROM golang:${GO_VERSION}-alpine AS builder

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

WORKDIR /src

# Fetch dependencies to cache before next build
COPY ./go.mod ./go.sum ./
RUN go mod download

# import source code
COPY ./ ./ 

# build statically linked binary to /app
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app .

FROM scratch AS final

COPY --from=builder /app/ app

EXPOSE 8080

ENTRYPOINT ["/app"]

