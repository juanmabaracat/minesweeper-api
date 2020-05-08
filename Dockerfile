FROM golang:1.14

RUN mkdir /app
WORKDIR /app

# Copy and download dependency using go mod
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o minesweeper-api ./src

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /app/minesweeper-api .

# Export necessary port
EXPOSE 8080

# Command to run when starting the container
CMD ["/dist/minesweeper-api"]
