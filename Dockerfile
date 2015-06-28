FROM debian:latest

EXPOSE 80

ADD config.toml /config.toml
ADD goji-skeleton /goji-skeleton
#ADD node_modules /node_modules
#ADD static /static
#ADD templates /templates

CMD ["/goji-skeleton"]
