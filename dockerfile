# # Build executable stage
# FROM golang

# # Add Maintainer Info
# LABEL maintainer="Nalin <nalin12894@gmail.com>"

# # Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy go mod and sum files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download
# # Copy the source from the current directory to the Working Directory inside the container
# COPY *.go ./
# COPY ./public ./public
# COPY gqldb.db ./
# # Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# # Build final image
# FROM alpine:latest

# RUN apk --no-cache add ca-certificates
# COPY --from=0 /app .

# # Command to run
# ENTRYPOINT ["/main"]



#
# builder image
#
FROM golang:1.13-alpine3.11 as builder
RUN mkdir -p /build/src

# make sure git and glide packages are present
RUN apk update \
    && apk add --update git \
    && apk add --update openssh

# standard golang env setup
ENV GOPATH=/build
ENV GOBIN=/usr/local/go/bin
# enable module support (package management)
ENV GO111MODULE=on
WORKDIR $GOPATH/src

# get main project from git
RUN git clone https://github.com/Nikhil12894/gqlwithgo.git \
    && ls /build/src/gqlwithgo
WORKDIR $GOPATH/src/gqlwithgo

# compile, place executable into /build
# by default, use git HEAD of external package "go-myutil"
ARG BRANCH=HEAD
RUN go get github.com/fabianlee/go-myutil@$BRANCH \
    && go list -m all \
    && CGO_ENABLED=0 GOOS=linux go build -a -o out . \
    && cp out $GOPATH/.
# COPY ./public /build/out
# COPY  gqldb.db /build/out
# intermediate executable
CMD [ "/build/out" ]

#
# generate clean, final image for end users
#
FROM alpine:3.11.3
# copy golang binary into container
COPY --from=builder /build/out .
COPY ./public ./public
COPY config.json .
# executable
CMD [ "./out" ]
