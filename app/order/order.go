package order

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"cleanarchitecture/app/usersession"
	"cleanarchitecture/biz/event"
	"cleanarchitecture/biz/order"
)

type Input struct {
	OrderDetails map[string]uint // 商品ID、数量
}

type Output struct {
	OrderID string
}

// OrderNow is usecase where a user confirmed an order.
func OrderNow(ctx context.Context, input Input) (*Output, error) {
	fmt.Println("ユーザーが注文を確定するユースケース")

	session, ok := usersession.FromContext(ctx)
	if !ok {
		return nil, errors.New("session not found")
	}

	for k, v := range input.OrderDetails {
		fmt.Println(k, v)
	}

	o := order.Order{
		ID:      uuid.New().String(),
		Details: input.OrderDetails,
		UserID:  session.UserID,
	}

	if err := order.Validate(ctx, o); err != nil { // 注文内容のチェック
		return nil, err
	}

	event.ConfirmedOrder.Publish(ctx, o)

	return &Output{OrderID: o.ID}, nil // 同期（request/response） であれば戻り値で結果を返す
}
