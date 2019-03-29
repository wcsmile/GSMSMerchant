package util

import (
	"coupon/customer/recv-coupon/modules/common"
	"fmt"
)

func SendMessageQueue(queueData map[string]string) {

	if queueData == nil {
		return
	}
	container := common.Container
	queueObj, err := container.GetQueue("queue")
	if err != nil {
		fmt.Println("SendMessageQueue:", err)
		return
	}
	for k, v := range queueData {
		fmt.Println("SendMessageQueue:", k, v)

		err = queueObj.Push(k, v)
		fmt.Println("SendMessageQueue:", err)
	}
}
