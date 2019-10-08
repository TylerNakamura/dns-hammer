FROM golang:latest

LABEL maintainer="Tyler Nakamura dns-hammer@tylernakamura.com"

WORKDIR /

ADD dnshammer /
ADD domains.txt /

CMD ["/dnshammer"]
