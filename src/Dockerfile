# Start from base image 1.12.13:
FROM golang:1.16.2

ENV LOG_LEVEL=info

# Configure the repo url so we can configure our work directory:
ENV REPO_URL=github.com/tony-landreth/bookstore_items-api

# Setup out $GOPATH
ENV GOPATH=/app

ENV APP_PATH=$GOPATH/src/$REPO_URL

# Copy the entire source code from the current directory to $WORKPATH
ENV WORKPATH=$APP_PATH/src
WORKDIR $WORKPATH
COPY . .

RUN go build -o items-api .

# Expose port 8080 to the world:
EXPOSE 8080

CMD ["./items-api"]