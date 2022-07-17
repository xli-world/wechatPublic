package flyweight

func main() {
	shop := newTeaShop()
	shop.takeOrder("redTea", 1)
	shop.takeOrder("redTea", 2)
	shop.takeOrder("greenTea", 5)
}
