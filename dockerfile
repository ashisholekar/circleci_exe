FROM golang:alpine
WORKDIR /build
COPY abc.go .
RUN go build -o abc abc.go
CMD ["./abc"]