# 소개

upbit-go 는 [upbit api](https://docs.upbit.com/)를 이용한 라이브러리 패키지 입니다. 

내부적으로 jwt 는 "github.com/dgrijalva/jwt-go", uuid 는 "github.com/google/uuid" 패키지를 사용합니다.

또한 요청 수 제한에 대한 처리는 하지 않고 API 마다 Remaining 구조체를 반환합니다. 개발시 참고 바랍니다.


### 설치

```bash
go get -u github.com/sangx2/upbit
```

## Getting started

더 많은 테스트 예문은 _test 이름을 가진 파일들을 참고 바랍니다.

### Exchange API

```go
package main

import (
	"fmt"

	"github.com/sangx2/upbit"
)

func main() {
	u := upbit.NewUpbit("AccessKey", "SecretKey")

	accounts, remaining, e := u.GetAccounts()
	if e != nil {
		fmt.Println("GetAccounts error : %s", e.Error())
	} else {
		fmt.Printf("GetAccounts[remaining:%+v]\n", *remaining)
		for _, account := range accounts {
			fmt.Printf("%+v\n", *account)
		}
	}
}
```

### Quotation API

Quotation API 는 "AccessKey"와 "SecretKey"가 필요하지 않습니다.

```go
package main

import (
	"fmt"

	"github.com/sangx2/upbit"
)

func main() {
    u := upbit.NewUpbit("", "")

	markets, remaining, e := u.GetMarkets()
	if e != nil {
		fmt.Println("GetMarkets error : %s", e.Error())
	} else {
		fmt.Printf("GetMarkets[remaining:%+v]\n", *remaining)
		for _, market := range markets {
			fmt.Printf("%+v\n", *market)
		}
	}
}
```