FROM ubuntu:15.10

# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		ca-certificates \
		curl \
		git \
		make \
		build-essential \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.6
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 5470eac05d273c74ff8bac7bef5bad0b5abbd1c4052efbdbc8db45332e836b0b

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

WORKDIR /go/src/github.com/gotravelydia/platform
ADD . /go/src/github.com/gotravelydia/platform

RUN go get -u github.com/tools/godep
RUN pwd
RUN ls

RUN godep restore -v
RUN go install ./...

# Removed unnecessary packages
#RUN apt-get autoremove -y

# Clear package repository cache
#RUN apt-get clean all

ENTRYPOINT /go/bin/runner
# application listening port
EXPOSE 8100
