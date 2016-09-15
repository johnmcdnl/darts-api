FROM golang:1.7
ADD . /go/src/github.com/johnmcdnl/darts
RUN go install github.com/johnmcdnl/darts
ENTRYPOINT /go/bin/darts
EXPOSE 4500