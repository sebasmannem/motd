build: build_helloworld build_cli

build_helloworld:
	go build ./cmd/helloworld -o ./bin/helloworld

build_cli:
	go build ./cmd/motd-cli - ./bin/cli

run: run_cli run_helloworld

run_cli:
	./bin/cli

run_helloworld:
	./bin/helloworld
