# luhn
クレジットカード番号 生成/検証

## 使い方

### クレジットカード番号の生成

```go
package main

import (
	"github.com/goccha/luhn"
)

func main() {
	// 16桁のクレジットカード番号を生成
	ccn := luhn.New().Generate(16)

	luhn.Verify(ccn)
}
```