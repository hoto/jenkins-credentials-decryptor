[![Docker hub](https://images.microbadger.com/badges/image/hoto/jenkins-credentials-decryptor.svg)](https://microbadger.com/images/hoto/jenkins-credentials-decryptor "Get your own image badge on microbadger.com")
[![CircleCI](https://circleci.com/gh/hoto/jenkins-credentials-decryptor/tree/master.svg?style=svg)](https://circleci.com/gh/hoto/jenkins-credentials-decryptor/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/jenkins-credentials-decryptor)](https://goreportcard.com/report/github.com/hoto/jenkins-credentials-decryptor)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/jenkins-credentials-decryptor/maintainability)
[![Release](https://img.shields.io/github/release/hoto/jenkins-credentials-decryptor.svg?style=flat-square)](https://github.com/hoto/jenkins-credentials-decryptor/releases/latest)
# Jenkins Credentials Decryptor

Command line tool for decrypting and dumping Jenkins credentials.

### What is this all about

Jenkins stores encrypted credentials in `credentials.xml` file.  
To decrypt them you need the `master.key` and `hudson.util.Secret` files.  

All three files are located inside Jenkins home directory:

    $JENKINS_HOME/credentials.xml 
    $JENKINS_HOME/secrets/master.key
    $JENKINS_HOME/secrets/hudson.util.Secret

### Run using a binary

Download binary from [releases](https://github.com/hoto/jenkins-credentials-decryptor/releases), Linux and Mac only:

    curl -L \
      "https://github.com/hoto/jenkins-credentials-decryptor/releases/download/0.0.5-alpha/jenkins-credentials-decryptor_0.0.5-alpha_$(uname -s)_$(uname -m)" \
       -o jenkins-credentials-decryptor

    chmod +x jenkins-credentials-decryptor

SSH into Jenkins box and run:

    ./jenkins-credentials-decryptor \
      -m $JENKINS_HOME/secrets/master.key \
      -s $JENKINS_HOME/secrets/hudson.util.Secret \
      -c $JENKINS_HOME/credentials.xml 
      
Or if you have the files locally:

    ./jenkins-credentials-decryptor \
      -m master.key \
      -s hudson.util.Secret \
      -c credentials.xml 
      
### Run using docker
    
If you are worried about me sending your credentials over the network (I can assure you I don't do that) 
then run a container with disabled network:

From Jenkins box:

    docker run \
      --rm \
      --network none \
      --workdir / \
      --mount "type=bind,src=$JENKINS_HOME/secrets/master.key,dst=/master.key" \
      --mount "type=bind,src=$JENKINS_HOME/secrets/hudson.util.Secret,dst=/hudson.util.Secret" \
      --mount "type=bind,src=$JENKINS_HOME/credentials.xml,dst=/credentials.xml" \
      docker.io/hoto/jenkins-credentials-decryptor:latest \
      /jenkins-credentials-decryptor \
        -m master.key \
        -s hudson.util.Secret \
        -c credentials.xml 

With files locally:

    docker run \
      --rm \
      --network none \
      --workdir / \
      --mount "type=bind,src=$PWD/master.key,dst=/master.key" \
      --mount "type=bind,src=$PWD/hudson.util.Secret,dst=/hudson.util.Secret" \
      --mount "type=bind,src=$PWD/credentials.xml,dst=/credentials.xml" \
      docker.io/hoto/jenkins-credentials-decryptor:latest \
      /jenkins-credentials-decryptor \
        -m master.key \
        -s hudson.util.Secret \
        -c credentials.xml 
        
### Build the binary yourself

If you are worried about executing a random binary from the internet then:

    git clone https://github.com/hoto/jenkins-credentials-decryptor.git
    make build
    
Binary will be in the `bin` folder.

---
 
### Development

Get:

    go get github.com/hoto/jenkins-credentials-decryptor/cmd/jenkins-credentials-decryptor/

Download dependencies:

    make dependencies

Build and test:

    make clean
    make build
    make test

Install to global golang bin directory:

    make install

---
_Following_ [_Standard Go Project Layout_](https://github.com/golang-standards/project-layout)
