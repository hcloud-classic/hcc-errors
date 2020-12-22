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

### version 수정 규칙
1. go module의 버전 명명법을 따른다. `v[major].[minor].[patch]`
2. 단순 에러 코드 업데이트는 patch 버전을 증가시킨다.
3. 새로운 리시버 함수 추가시 minor 버전을 증가시키고 patch 버전을 0으로 초기화 한다.
4. hcc_error의 전체적인 구조변경이 있을 때 major 버전을 증가시킨다.
5. minor, major 증가를 요구하는 변경은 Discussion을 통해 필요성을 검증 한다.
6. minor, major 버전 증가 후에는 `go get`을 
