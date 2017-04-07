BINARY=go-ip-api

VERSION=1.1.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BUILD=${BUILD}"

default:
	go build ${LDFLAGS} -o ${BINARY}
	./md2html.pl

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY}; fi
