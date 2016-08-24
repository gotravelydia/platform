FROM ubuntu:15.10

# Install GCC.
RUN apt-get update
RUN apt-get install -y --no-install-recommends \
	g++ \
	gcc \
	libc6-dev \
	ca-certificates \
	curl \
	git \
	make \
	build-essential

# Clear apt-get cache, it could get big.
RUN rm -rf /var/lib/apt/lists/*

# GoLang version and download path.
ENV GOLANG_VERSION 1.6
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 5470eac05d273c74ff8bac7bef5bad0b5abbd1c4052efbdbc8db45332e836b0b

# Download GoLang.
RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz

# Verify the tar.
RUN echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c -

# Unzip.
RUN tar -C /usr/local -xzf golang.tar.gz

# Delete the tar.
RUN rm golang.tar.gz

# Set the Go path.
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# Set the workspace.
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
WORKDIR /go/src/github.com/gotravelydia/platform

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/gotravelydia/platform

# Install Godep dependency manager.
RUN go get -u github.com/tools/godep

# Install all the dependencies defined in Godeps.json.
RUN go get github.com/gorilla/context
RUN godep restore

# Build the platform binary inside the container.
RUN go install ./...

# Start running platform.
ENTRYPOINT /go/bin/runner

# Document that the service listens on port 8100.
EXPOSE 8100
