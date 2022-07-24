// 场景描述：
// 在电商网站中，商品时不时地会出现缺货情况。可能会有客户对于缺货的特定商品表现出兴趣。客户可以订阅其感兴趣的特定商品，商品可用时便会收到通知。同时，多名客户也可订阅同一款产品。
package observer

func main() {
	shirtItem := newItem("Nike Shirt")

	observerFirst := &customer{id: "abc@qq.com"}
	observerSecond := &customer{id: "xyz@qq.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
