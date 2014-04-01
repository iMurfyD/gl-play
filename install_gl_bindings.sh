#! /bin/bash
#currently for mac os x with glfw installed in /opt/local (macports) ONLY!!!!
export GOPATH=.
go get github.com/go-gl/gl
cd src/github.com/go-gl/gl
CGO_CFLAGS="-I/opt/local/include/" CGO_LDFLAGS="-L/opt/local/lib/" go install
 go get github.com/go-gl/glfw3
 cd src/github.com/go-gl/glfw3
 CGO_CFLAGS="-I/opt/local/include/" CGO_LDFLAGS="-L/opt/local/lib/" go install
