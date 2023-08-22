# Use an official Go runtime as a parent image
FROM golang:1.19.1
LABEL authors="Xun Yang"

# Install necessary dependencies
RUN apt-get update && apt-get install -y wget gnupg

# Install Google Chrome
RUN wget -q -O - https://dl.google.com/linux/linux_signing_key.pub | apt-key add -
RUN echo "deb [arch=amd64] https://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list
RUN apt-get update && apt-get install -y google-chrome-stable

# Set the working directory in the container
WORKDIR /webpagefetcher

# Copy the local package files to the container's workspace
COPY . /webpagefetcher

# Build the Go application inside the container
RUN go mod tidy

RUN go build -o webpagefetcher ./cmd/webpagefetcher/main.go

# Run the compiled binary
CMD ["./webpagefetcher"]