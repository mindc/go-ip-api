BINARY=fasthttp

VERSION=1.2.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.BUILD=${BUILD}"

default: 
	go build ${LDFLAGS}
	test -x ./md2html.pl && ./md2html.pl

install:
	go install ${LDFLAGS}
	test -x ./md2html.pl && ./md2html.pl

clean:
	test -f ${BINARY} && rm -v ${BINARY}

test: 
	go test -v
