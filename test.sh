go test -v ./stack_go
go build -o stack ./stack_c
./stack
rm -rf ./stack
