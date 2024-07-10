# Specifies a parent image
FROM golang:1.21
 
# Creates an app directory to hold your app’s source code
WORKDIR /app
 
# Copies everything from your root directory into /app
COPY . .
 
RUN go run github.com/steebchen/prisma-client-go generate --schema ./prisma


# Installs Go dependencies
RUN go mod tidy
 
RUN go run github.com/steebchen/prisma-client-go generate --schema ./prisma

# Builds your app with optional configuration
RUN go build ./cmd/server
 
# Tells Docker which network port your container listens on
EXPOSE 8080
EXPOSE 5000


RUN chmod a+x ./server
# Specifies the executable command that runs when the container starts
CMD [ "./server" ]
