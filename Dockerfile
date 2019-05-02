FROM golang:1.12.4 AS builder
ENV PROJECT_PATH="$GOPATH/src/github.com/hoto/jenkins-credentials-decryptor"
ENV BINARY_PATH="$PROJECT_PATH/bin/jenkins-credentials-decryptor"
WORKDIR $PROJECT_PATH
COPY . .
RUN make build

FROM scratch
ENV PROJECT_PATH="$GOPATH/src/github.com/hoto/jenkins-credentials-decryptor"
ENV BINARY_PATH="$PROJECT_PATH/bin/jenkins-credentials-decryptor"
COPY --from=builder $BINARY_PATH /
CMD [". /jenkins-credentials-decryptor"]
