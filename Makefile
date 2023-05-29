build: build_helloworld build_cli

build_helloworld:
	go build ./cmd/helloworld

build_cli:
	go build ./cmd/motd-cli

run: run_cli run_helloworld

run_cli:
	./motd-cli

run_helloworld:
	./helloworld
