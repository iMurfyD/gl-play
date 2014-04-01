This is just me playing around with opengl in golang. If anyone cares, I am using thsee bindings for OpenGL:
https://www.github.com/go-gl/gl
https://www.github.com/go-gl/glfw3

To Compile:
**note I am using OS X Mavericks with glew installed via macports in /opt/local, this will not work under any other condition
1. Clone Directory
2. Execute install_gl_bindings.sh
3. Set GOPATH to root directory and use "go install main"
4. The binary will be installed in GOPATH/bin
