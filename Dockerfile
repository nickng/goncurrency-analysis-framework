FROM nickng/godel:latest
MAINTAINER Nicholas Ng <nickng@nickng.io>
WORKDIR /root
ENV GOPATH /go
ENV GOLANG_VERSION 1.8.7
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN apt-get update -y && \
	apt-get install -y --no-install-recommends wget ca-certificates git gcc libncurses5; \
	wget -O go.tgz "https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz"; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	\
	apt-get remove wget ca-certificates -y && \
	apt-get autoremove -y --purge && \
	rm -rf /var/lib/apt/lists/*
RUN go version
RUN go get -v -d github.com/nickng/gospal/... && \
	go get -v -d go.uber.org/zap/... && cd /go/src/go.uber.org/zap && git checkout 06a239 && \
	go get -v -d github.com/nickng/goncurrency-analysis-framework/... && \
	cd /go/src/github.com/nickng/goncurrency-analysis-framework
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN go install -v github.com/nickng/gospal/cmd/migoinfer
RUN go install -tags native -v github.com/nickng/goncurrency-analysis-framework/cmd/demotool
CMD [ "demotool" ]
