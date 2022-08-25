package payment

import (
	"context"
	"fmt"

	"cleanarchitecture/biz/order"
)

func Pay(ctx context.Context, order order.Order) {
	fmt.Println("◇◇◇ 指定された決済方法で支払処理を実行します")
}
