FROM golang:1.14.6

# Set the Current Working Directory inside the container
WORKDIR /app/url-shortener

# Copy everything to the PWD inside the container
COPY . .

# Download all dependencies
RUN go get -d -v ./...

RUN go build -o ./out/url-shortener .

EXPOSE 8080

# Run the executable
CMD ["./out/url-shortener"]
