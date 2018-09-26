FROM golang:1.9.1

WORKDIR /go/src/github.com/mlvhub/learning-go/grpc-kube/gcd
COPY gcd .
COPY pb ../pb

RUN go get -v ./...
RUN go install -v ./...

EXPOSE 3000

CMD [ "gcd" ]


