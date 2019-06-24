# 소개

한국인이 대부분 이용할 것 같아서 한글로 남깁니다.

upbit-go 는 [upbit api](https://docs.upbit.com/)를 이용한 라이브러리 패키지 입니다. 

내부적으로 jwt 는 "github.com/dgrijalva/jwt-go", uuid 는 "github.com/google/uuid" 패키지를 사용합니다.

또한 요청 수 제한에 대한 처리는 하지 않고 API 마다 Remaining 구조체를 반환합니다. 개발시 참고하시기 바랍니다.


### 설치

```bash
go get -u github.com/sangx2/upbit-go
```

## Getting started

### Exchange API

더 많은 테스트 예문은 upbit_test.go 파일의 TestExchange 함수를 참고 하시기 바랍니다.

```go
package main

import (
	"fmt"

	upbit "github.com/sangx2/upbit-go"
)

func main() {
	u := upbit.NewUpbit("AccessKey", "SecretKey")

	accounts, remaining, e := u.GetAccounts()
	if e != nil {
		fmt.Println("GetAccounts error : %s", e)
	} else {
		fmt.Printf("GetAccounts[remaining:%+v]", *remaining)
		for _, account := range accounts {
			fmt.Printf("%+v", *account)
		}
	}
}
```

### Quotation API

Quotation API 는 "AccessKey"와 "SecretKey"가 필요하지 않습니다.

더 많은 테스트 예문은 upbit_test.go 파일의 TestQuotation 함수를 참고 하시기 바랍니다.

```go
package main

import (
	"fmt"

	upbit "github.com/sangx2/upbit-go"
)

func main() {
    u := NewUpbit("", "")

	markets, remaining, e := u.GetMarkets()
	if e != nil || len(markets) == 0 {
		fmt.Println("GetMarkets error : %s", e)
	} else {
		fmt.Printf("GetMarkets[remaining:%+v]", *remaining)
		for _, market := range markets {
			fmt.Printf("%+v", *market)
		}
	}
}
```