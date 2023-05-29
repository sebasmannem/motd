build: build_helloworld

build_helloworld:
	go build ./cmd/helloworld

run: run_helloworld

run_helloworld:
	./helloworld
