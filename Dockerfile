FROM golang:1.14.6

# Set the Current Working Directory inside the container
WORKDIR /app/url-shortener

# Get dependencies
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything to the PWD inside the container
COPY . .

# Build via make
RUN make build

EXPOSE 8080

# Run the executable
CMD ["./dist/url-shortener"]
