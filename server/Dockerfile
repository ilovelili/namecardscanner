FROM golang:1.8.3-alpine

# alpine do not include git so add it
RUN apk add --no-cache git

# Set apps working directory
WORKDIR /app

# Set an env var that matches $GOPATH
ENV SRC_DIR=/go/src

# Copy the local package files to the container's workspace
ADD /src/. $SRC_DIR/

# get go dependency
RUN echo "Installing Go dependencies ..."
RUN go get github.com/gorilla/mux
RUN go get github.com/rs/cors
RUN go get golang.org/x/net/context
RUN go get golang.org/x/oauth2/google
RUN go get google.golang.org/api/vision/...

# Test it, build it and copy to compiled directory
RUN cd $SRC_DIR/namecardscanner; go build -o namecardscanner; cp -r namecardscanner serviceaccount.json download /app/;

ENTRYPOINT ["./namecardscanner"]

# take 3000
EXPOSE 3000