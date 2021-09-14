FROM golang:latest
MAINTAINER matheuscoding@gmail.com
RUN mkdir /gamesort
COPY . /gamesort
WORKDIR /gamesort/src
RUN go mod tidy
CMD go run *.go
