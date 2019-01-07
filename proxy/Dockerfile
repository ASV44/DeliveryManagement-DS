############################
# STEP 1 build executable binary
############################

FROM golang:alpine AS builder

# Set environment variables

ENV APP DeliveryManagement-DS
ENV SERVICE proxy
ENV WORKDIR ${GOPATH}/src/github.com/ASV44/${APP}

# Install git.
# Git is required for fetching the dependencies.

RUN apk update && apk add --no-cache git

#Copy source code to build image

COPY cmd/$SERVICE $WORKDIR/cmd/$SERVICE
COPY $SERVICE $WORKDIR/$SERVICE/

#Copy dependecies list files

COPY $SERVICE/Gopkg.toml $WORKDIR/
COPY $SERVICE/Gopkg.lock $WORKDIR/

#Change container work directory

WORKDIR $WORKDIR

# Fetch dependencies.
# Using dep.

RUN set -x && \
    go get github.com/golang/dep/cmd/dep && \
    dep ensure -v

# Build the binary.

RUN go build -v -o /go/bin/main cmd/$SERVICE/main.go

###############################################
# STEP 2 build a small image, production build
##############################################

# FROM alpine

# Copy our static executable.

# COPY --from=builder /go/bin/main /go/bin/main

ENV PORT=8000
EXPOSE $PORT


# Run the binary.
ENTRYPOINT ["/go/bin/main"]
