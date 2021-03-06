FROM --platform=linux/amd64 golang:alpine

# RUN useradd -u 8877 whm
# USER whm

COPY --from=docker/compose:1.29.2 /usr/local/bin/docker-compose /usr/bin/docker-compose

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Export necessary port
EXPOSE 3000

# Command to run when starting the container
CMD ["/dist/main"]