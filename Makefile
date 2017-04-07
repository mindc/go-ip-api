BINARY=go-ip-api

VERSION=1.1.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BUILD=${BUILD}"

default:
	go build ${LDFLAGS} -o ${BINARY}
	test -x ./md2html.pl && ./md2html.pl

clean:
	test -f ${BINARY} && rm -v ${BINARY}
