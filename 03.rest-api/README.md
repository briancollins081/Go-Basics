# GO REST APIS

## Reference for testing
- Run a code coverage on all modules: `go test -v ./... -coverprofile cover.out`
- Convert a code coverage output into html for viewing in the browser using go tool: `go tool cover -html=cover.out -o cover.html`

## Reference for Benchmarks
- Run a benchmark by:
  1. `cd module`  
  2. `go test -bench .`  