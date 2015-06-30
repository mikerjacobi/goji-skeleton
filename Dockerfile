FROM golang:latest

EXPOSE 80

ADD config.toml /config.toml
#ADD . /go/src/
#ADD node_modules /node_modules
#ADD static /static
#ADD templates /templates
CMD go build /go/src/
CMD ["/go/src/goji-skeleton"]
