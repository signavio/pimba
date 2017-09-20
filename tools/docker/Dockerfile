FROM golang:1.9 as builder

WORKDIR /go/src/github.com/signavio/pimba
ADD . .
RUN go get && go install

FROM scratch
COPY --from=builder /go/bin/pimba .

ENTRYPOINT ["/pimba"]
EXPOSE 8080
CMD ["api", "--storage", "/var/www", "--secret", "$JWT_SECRET"]
