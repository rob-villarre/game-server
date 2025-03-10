# Specifies a parent image
FROM golang:1.24.0-alpine

# Install necessary build tools for CGO and GLFW
# RUN apk add --no-cache \
#   gcc \
#   g++ \
#   make \
#   musl-dev \
#   linux-headers \
#   libx11-dev \
#   libxcursor-dev \
#   libxrandr-dev \
#   libxinerama-dev \
#   libxi-dev \
#   mesa-dev \
#   wayland-dev

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY /go.mod /go.sum ./

# Download all dependancies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY ./cmd/server ./cmd/server
COPY ./internal ./internal
 
# Builds your app with optional configuration
# RUN CGO_ENABLED=1 GOOS=linux go build -o /bin/server ./cmd/server
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/server ./cmd/server
 
# Tells Docker which network port your container listens on
EXPOSE 3000/udp
 
# Specifies the executable command that runs when the container starts
CMD [ "/bin/server", "3000" ]