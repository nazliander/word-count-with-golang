FROM golang:1.14

# Setup working dir
RUN mkdir /app 
ADD ./app /app
WORKDIR /app

# Download all the dependencies
RUN go get -d -v ./...

# Build
RUN go build *.go

# This container exposes port 7979
EXPOSE 7979

# Run the executable
CMD ["./bookAnalytics"]