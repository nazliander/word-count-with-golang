FROM golang:1.14

# Setup working dir
ADD ./app /app
WORKDIR /app

# Build
RUN go build

# This container exposes port 7979
EXPOSE 7979

# Run the executable
CMD ["./app"]
