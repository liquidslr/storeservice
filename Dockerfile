# iron/go:dev is the alpine image with the go tools added
FROM golang:1.11 

WORKDIR /app

# Set an env var that matches the github repo name
ENV SRC_DIR=/go/src/github.com/liquidslr/storeservice/

# Add the source code:
ADD . $SRC_DIR

# Build it:
RUN cd $SRC_DIR; go build -o myapp; cp myapp /app/;


ENTRYPOINT ["./myapp"]