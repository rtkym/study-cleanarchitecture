package order

import (
	"context"
	"fmt"
)

type Order struct {
	ID      string          // 注文番号
	Details map[string]uint // 商品ID、数量
	UserID  string          // 注文者
}

func Validate(ctx context.Context, o Order) error {
	fmt.Println("★★★ 注文内容を検査する処理")
	return nil
}
