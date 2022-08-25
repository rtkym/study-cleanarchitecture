package example

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"cleanarchitecture/app/order"
	"cleanarchitecture/app/usersession"
)

func UserConfirmOrder(ctx context.Context, req http.Request) {
	session := usersession.UserSession{UserID: "bob"}
	ctx = usersession.NewContext(ctx, &session)

	// HTTPのデータ構造をユースケースの入力データに変換する
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	orderItems := new(map[string]uint)
	if err := json.Unmarshal(body, orderItems); err != nil {
		log.Fatalln(err)
	}

	input := order.Input{
		OrderDetails: *orderItems,
	}

	//　ユースケースの実行
	output, err := order.OrderNow(ctx, input)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}

	// 本来ならHTTPレスポンスを返す
}
