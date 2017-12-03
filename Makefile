all: main.go 
	go build -o ./dist/server.exe

clean: 
	$(RM) ./dist/server.exe

install:
	go get github.com/gorilla/mux
	go get github.com/ttacon/chalk

run: 
	./dist/server.exe