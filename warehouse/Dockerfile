############################
# STEP 1 build executable binary
############################

FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.

ENV SERVICE_NAME DeliveryManagement-DS
ENV APP /src/${SERVICE_NAME}/
ENV WORKDIR ${GOPATH}${APP}

RUN apk update && apk add --no-cache git
COPY . WORKDIR
WORKDIR WORKDIR

# Fetch dependencies.
# Using go get.

RUN go get -d -v

# Build the binary.

RUN go build -o /go/bin/main

############################
# STEP 2 build a small image
############################

FROM alpine

# Copy our static executable.

COPY --from=builder /go/bin/main /go/bin/main

ENV PORT=8080
EXPOSE $PORT

# Run the hello binary.
ENTRYPOINT ["/go/bin/main"]
