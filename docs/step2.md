# Step 2
Building hello world

## Step 2.1
Intialize
```
mkdir ~/git/sebasmannem/motd
cd ~/git/sebasmannem/motd
git init
go mod init github.com/sebasmannem/motd
mkdir -p {cli,pkg,internal}
```

## Step 2.2
Add code (hello world)
```
mkdir cli/helloworld
vim cli/helloworld/main.go
```

## step 2.3
```
go run ./cmd/helloworld
```

### step 2.4
build it and run it as binary
```
go build ./cmd/helloworld
./helloworld
```
