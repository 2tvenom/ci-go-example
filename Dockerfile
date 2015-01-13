FROM google/golang:latest
MAINTAINER Pavel <Ven> Gulbin <2tvenom@gmail.com>

WORKDIR $GOPATH/src/github.com/2tvenom/ci-php-example

EXPOSE 8080

ADD . $GOPATH/src/github.com/2tvenom/ci-php-example
RUN go install github.com/2tvenom/ci-php-example

ENTRYPOINT ["ci-php-example"]