FROM golang:1.10

WORKDIR /go/src/fortune-server
COPY *.go .
RUN go get -d -v ./...
RUN go install -v ./...

RUN apt-get -y update && apt-get install -y fortunes

EXPOSE 80
CMD ["fortune-server", "-port=80"]
