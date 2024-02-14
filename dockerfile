#declaring our go version
FROM golang:1.19.0
#our container is a unix container
WORKDIR /usr/src/app
#Below command is for hot reloading
RUN go install github.com/cosmtrek/air@latest
#copy contents into the container
COPY . .
#tidy up the dependencies
RUN go mod tidy