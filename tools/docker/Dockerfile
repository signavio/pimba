FROM golang:1.9-alpine

ADD . /go/src/github.com/signavio/pimba
RUN cd /go/src/github.com/signavio/pimba && go get && go install
ENTRYPOINT ["/go/bin/pimba"]
EXPOSE 8080
CMD ["api", "--storage", "/var/www", "--secret", "$JWT_SECRET"]
