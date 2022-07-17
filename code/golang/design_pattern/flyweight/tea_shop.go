package flyweight

import "fmt"

type teaShop struct {
	orders map[int]tea
}

func (t *teaShop) takeOrder(teaType string, tableId int) {
	tea, _ := getTeaFactorySingleInstance().getTeaByType(teaType)
	t.orders[tableId] = tea
}

func (t *teaShop) serve() {
	for tableId, tea := range t.orders {
		fmt.Printf("Serving %s to table # %v", tea.getName(), tableId)
	}
}

func newTeaShop() *teaShop {
	return &teaShop{orders: make(map[int]tea)}
}
