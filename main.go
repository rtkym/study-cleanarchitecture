package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"cleanarchitecture/biz/event"
	"cleanarchitecture/biz/payment"
	"cleanarchitecture/biz/picking"
	"cleanarchitecture/ifadapter/example"
)

func main() {
	event.ConfirmedOrder.Subscribe(picking.Picking)
	event.ConfirmedOrder.Subscribe(payment.Pay)

	ctx := context.Background()

	req := http.Request{
		Body: ioutil.NopCloser(bytes.NewReader([]byte(`{
			"book": 3,
			"food": 8
		}`))),
	}

	example.UserConfirmOrder(ctx, req)

	event.ConfirmedOrder.Wait(ctx)
}
