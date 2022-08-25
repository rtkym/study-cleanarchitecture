package picking

import (
	"context"
	"fmt"

	"cleanarchitecture/biz/order"
)

func Picking(ctx context.Context, order order.Order) {
	fmt.Println("◆◆◆ 注文された商品をピッキングする処理")
}
