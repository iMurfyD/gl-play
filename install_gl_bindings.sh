#! /bin/bash
#currently for mac os x with glfw installed in /opt/local (macports) ONLY!!!!
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export GOPATH=$DIR
echo $GOPATH
rm -r $GOPATH/src/github.com/go-gl
echo "Getting gl"
go get github.com/go-gl/gl
cd src/github.com/go-gl/gl
echo "Installing gl"
CGO_CFLAGS="-I/opt/local/include/" CGO_LDFLAGS="-L/opt/local/lib/" go install
echo "Getting glfw3"
go get github.com/go-gl/glfw3
cd src/github.com/go-gl/glfw3
echo "Installing glfw3"
CGO_CFLAGS="-I/opt/local/include/" CGO_LDFLAGS="-L/opt/local/lib/" go install
