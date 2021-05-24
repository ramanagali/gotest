# steps to test

1. To run Tests
go test 

2. Verbose test output
go test -v

3. test coverage of go test
go test -cover

4. to generete coverprofile
go test -coverprofile=coverage.out

5. to view html coverprofile
go tool cover -html=coverage.out
