FROM golang:1.4

# Grab some Go tools
RUN go get golang.org/x/tools/cmd/cover \
    && go get golang.org/x/tools/cmd/vet \
    && go get github.com/smartystreets/goconvey

COPY . /go/src/github.com/vdemeester/miniature-todos
WORKDIR /go/src/github.com/vdemeester/miniature-todos

EXPOSE 9090

# Download dependencies
RUN go-wrapper download
# Install it
RUN go-wrapper install

ENTRYPOINT ["miniature-todos"]
