package event

import "cleanarchitecture/biz/order"

var ConfirmedOrder = newBroker[order.Order]("ConfirmedOrder")
