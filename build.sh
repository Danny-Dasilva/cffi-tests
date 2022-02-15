# go build -o awesome.so -buildmode=c-shared awesome.go

go build -o golang.so -buildmode=c-shared golang/*.go

go build -o test.so -buildmode=c-shared test.go
