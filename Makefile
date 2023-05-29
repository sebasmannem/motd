build: build_helloworld build_cli

build_helloworld:
	go build -o ./bin/helloworld ./cmd/helloworld

build_cli:
	go build -o ./bin/cli ./cmd/motd-cli

run: run_cli run_helloworld

run_cli:
	./bin/cli

run_helloworld:
	./bin/helloworld
