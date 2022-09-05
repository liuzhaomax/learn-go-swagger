SHELL = /bin/bash

pull:
	docker pull quay.io/goswagger/swagger

run:
	docker run --rm -it --env GOPATH=/go -v %CD%:/go/src -w /go/src quay.io/goswagger/swagger
