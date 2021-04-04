BINARY = bin/wfp-engine
GOARCH = amd64
# Build the project
build: clean windows
windows:
	CGO_ENABLED=1 CXX=x86_64-w64-mingw32-g++ CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -v -o ${BINARY}.exe main.go ;

clean:
	-rm -f ${BINARY}-*
