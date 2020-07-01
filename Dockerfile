FROM golang:1.14.4 AS builder
WORKDIR /go/src/github.com/hoto/jenkins-credentials-decryptor
COPY . .
RUN make build

FROM scratch
WORKDIR /
COPY --from=builder /go/src/github.com/hoto/jenkins-credentials-decryptor/bin/jenkins-credentials-decryptor .
CMD ["/jenkins-credentials-decryptor"]
