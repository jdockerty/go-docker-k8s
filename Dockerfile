# Latest Golang base image
FROM golang:latest as builder 

# Set working dir for container as /app
WORKDIR /app


# Copy from current dir to working dir inside of container
COPY . .

# Build Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# New stage, alpine latest image
FROM alpine:latest

# Installs certificates in image for SSL/TLS
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copies build from previous stage 
COPY --from=builder /app/* /app/html/* .

# Expose port 8080
EXPOSE 8080

# ./main can be used to run executable
CMD ["./main"]
