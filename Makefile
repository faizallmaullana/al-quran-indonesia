running: build run 

run:
	go run . 2

build: npm-install build-node build-go

build-node: 
	cd alQuranFrontEnd && npm run build

build-go:
	go build -o main .

npm-install:
	cd alQuranFrontEnd && npm install
