[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](/LICENSE)
[![Build status](https://github.com/hoto/jenkins-credentials-decryptor/workflows/Test/badge.svg?branch=master)](https://github.com/hoto/jenkins-credentials-decryptor/actions)
[![Release](https://img.shields.io/github/release/hoto/jenkins-credentials-decryptor.svg?style=flat-square)](https://github.com/hoto/jenkins-credentials-decryptor/releases/latest)
[![Powered By: goreleaser](https://img.shields.io/badge/powered%20by-goreleaser-green.svg?style=flat-square)](https://github.com/goreleaser/goreleaser)
[![Go Report Card](https://goreportcard.com/badge/github.com/hoto/jenkins-credentials-decryptor)](https://goreportcard.com/report/github.com/hoto/jenkins-credentials-decryptor)
[![Maintainability](https://api.codeclimate.com/v1/badges/27f61a82b9a5589f1a07/maintainability)](https://codeclimate.com/github/hoto/jenkins-credentials-decryptor/maintainability)
# Jenkins Credentials Decryptor

Command line tool for decrypting and dumping Jenkins credentials.

### What is this all about

Jenkins stores encrypted credentials in the `credentials.xml` file or in `config.xml`. 
To decrypt them you need the `master.key` and `hudson.util.Secret` files.  

All files are located inside Jenkins home directory:

    $JENKINS_HOME/credentials.xml 
    $JENKINS_HOME/secrets/master.key
    $JENKINS_HOME/secrets/hudson.util.Secret
    $JENKINS_HOME/jobs/example-folder/config.xml - Possible location

### Compatibility

I've tested this on Jenkins 1.625.1 and 2.141

### Run using a binary

Mac:

    brew install hoto/repo/jenkins-credentials-decryptor

Mac or Linux:

    curl -L \
      "https://github.com/hoto/jenkins-credentials-decryptor/releases/download/1.2.0/jenkins-credentials-decryptor_1.2.0_$(uname -s)_$(uname -m)" \
       -o jenkins-credentials-decryptor

    chmod +x jenkins-credentials-decryptor
    
Or manually download binary from [releases](https://github.com/hoto/jenkins-credentials-decryptor/releases).

Help:

    ./jenkins-credentials-decryptor --help
    ./jenkins-credentials-decryptor --version

SSH into Jenkins box and run:

    ./jenkins-credentials-decryptor \
      -m $JENKINS_HOME/secrets/master.key \
      -s $JENKINS_HOME/secrets/hudson.util.Secret \
      -c $JENKINS_HOME/credentials.xml \
      -o json
      
Or if you have the files locally:

    ./jenkins-credentials-decryptor \
      -m master.key \
      -s hudson.util.Secret \
      -c credentials.xml \
      -o json
      
### Run using docker
    
If you are worried about the binary sending your credentials over the network (it does not do that) 
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
        -c credentials.xml \
        -o json

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
        -c credentials.xml \
        -o json
        
### Build the binary yourself

If you are worried about executing a random binary from the internet then:

    git clone https://github.com/hoto/jenkins-credentials-decryptor.git
    make build
    
Binary will be located at `bin/jenkins-credentials-decryptor`.

---

### Example output

Json output format:

    $ ./jenkins-credentials-decryptor \
           -m master.key \
           -s hudson.util.Secret \
           -c credentials.xml \
           -o json
          
    [
      {
        "description": "Vault admin",
        "id": "vault-admin",
        "username": "admin",
        "password": "9cy7Mbw@1Omm7db@q6eP3k62Wm*ev#",
        "scope": "GLOBAL"
      }
    ]

Text output format:
 
    $ ./jenkins-credentials-decryptor \
           -m master.key \
           -s hudson.util.Secret \
           -c credentials.xml \
           -o text
          
    0
            description: Vault admin
            id: vault-admin
            username: admin
            password: 9cy7Mbw@1Omm7db@q6eP3k62Wm*ev#
            scope: GLOBAL

---
 
### Development

Clone:

    mkdir -p $GOPATH/src/github.com/hoto
    cd $GOPATH/src/github.com/hoto
    git clone https://github.com/hoto/jenkins-credentials-decryptor.git

Download dependencies:

    make dependencies

Build and test:

    make clean
    make build
    make test
    
Run a good ol' fashion manual smoke test:

    make smoke-test-json
    make smoke-test-text

Install to global golang bin directory:

    make install

---
_Following_ [_Standard Go Project Layout_](https://github.com/golang-standards/project-layout)
