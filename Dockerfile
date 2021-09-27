FROM golang:latest
MAINTAINER matheuscoding@gmail.com
COPY . /gamesort
WORKDIR /gamesort/src
RUN go mod tidy
CMD go run *.go
