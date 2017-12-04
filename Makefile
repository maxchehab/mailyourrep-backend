all: src/main.go
	cd src; go build -o ../dist/server.exe
	./dist/server.exe

build: src/main.go 	
	cd src; go build -o ../dist/server.exe

clean: 
	$(RM) ./dist/server.exe

install:
	go get github.com/gorilla/mux
	go get github.com/ttacon/chalk
	go get github.com/stripe/stripe-go

run: dist/server.exe
	./dist/server.exe