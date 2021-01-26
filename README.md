# hcc_errors
## 사용법
### Makefile 수정
```
build: ## Build the binary file
	@$(GOROOT)/bin/go get -u=patch github.com/hcloudclassic/hcc_errors
	@$(GOROOT)/bin/go build -o ${PROJECT_NAME} main.go
```

### version 수정
`errors.go`
```
const version = "1.0.1" // Change here after modify Error Code

/*** Match Enum squence with xxxList ***/
const (
        // code for MiddleWare
        modulename uint64 = (1 + iota) * 10000
        cello
```

