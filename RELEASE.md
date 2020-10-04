# How to release this app

### The quick way

Commit on master and push tag:

    git tag <breaking>.<feature>.<bugfix>
    git push --tags
    
### The safe way

Run smoke test to check if the binary is built correctly:

    make smoke-test-json
    make smoke-test-text

Build and test the binary produced by goreleaser:

    make goreleaser-dry-run-local
    
    ./dist/jenkins-credentials-decryptor_$(uname -s)_amd64/jenkins-credentials-decryptor --version
   
   
Git tag and push (instructions above). 