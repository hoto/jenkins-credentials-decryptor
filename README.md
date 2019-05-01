[![CircleCI](https://circleci.com/gh/hoto/jenkins-credentials-decryptor/tree/master.svg?style=svg)](https://circleci.com/gh/hoto/jenkins-credentials-decryptor/tree/master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/jenkins-credentials-decryptor)](https://goreportcard.com/report/github.com/hoto/jenkins-credentials-decryptor)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/jenkins-credentials-decryptor/maintainability)
[![Release](https://img.shields.io/github/release/hoto/jenkins-credentials-decryptor.svg?style=flat-square)](https://github.com/hoto/jenkins-credentials-decryptor/releases/latest)
# Jenkins Credentials Decryptor

Command line tool for decrypting and dumping Jenkins credentials.

### What is this all about

Jenkins stores encrypted credentials in `credentials.xml` file.  
To decrypt them you need the `master.key` and `hudson.util.Secret` files.  
All three files are located inside Jenkins home directory.

### Installation

Download binary from [releases](https://github.com/hoto/jenkins-credentials-decryptor/releases):

Linux and mac:

    curl -L \
      "https://github.com/hoto/jenkins-credentials-decryptor/releases/download/1.0.0/jenkins-credentials-decryptor_1.0.0_$(uname -s)_$(uname -m)" \
       -o jenkins-credentials-decryptor

    chmod +x jenkins-credentials-decryptor
    
### Usage

Ssh into the Jenkins box or copy the files locally then run:

    jenkins-credentials-decryptor \
      -m $JENKINS_HOME/secrets/master.key \
      -s $JENKINS_HOME/hudson.util.Secret \
      -c $JENKINS_HOME/credentials.xml 
    
### Development

Get:

    go get github.com/hoto/jenkins-credentials-decryptor/

Download dependencies:

    make dependencies

Build and test:

    make clean
    make build
    make test

Install to global golang bin directory:

    make install

