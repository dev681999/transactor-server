#########################
# STEP 1 compile a binary
#########################
FROM golang:1.23-bookworm AS builder

WORKDIR /app

COPY go.* ./

# download dependencies
RUN go mod download -x

# copy code
COPY cmd ./cmd
COPY pkg ./pkg

# Build the binary
RUN CGO_ENABLED=1 go build \
    -tags netgo,osusergo \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o server cmd/server/*.go

###################################
# STEP 2 compile healthcheck binary
###################################
FROM golang:1.23-bookworm AS healthcheck

WORKDIR /

# copy code
COPY healthcheck/ ./

# build binary
RUN CGO_ENABLED=1 go build \
    -tags netgo,osusergo \
    -ldflags='-w -s -extldflags "-static"' -a \
    -o healthcheck *.go

# we need the atlas binary for our migrations
FROM arigaio/atlas:0.28.1 AS atlas

############################
# STEP 3 build a final image
############################
FROM busybox:1.37

WORKDIR /app

# Copy healthcheck and atlas executable
COPY --from=healthcheck /healthcheck /bin/
COPY --from=atlas /atlas /bin/

# Copy our server executable
COPY --from=builder /app/server /bin

# copy base config
COPY config/base.yml ./
# copy migrations
COPY migrations /migrations

EXPOSE 8080

HEALTHCHECK \
    --interval=5s \
    --timeout=5s \
    CMD ["healthcheck", "http://localhost:8080/livez"]

# setup server with base config
ENTRYPOINT ["server", "-config", "/app/base.yml"]

