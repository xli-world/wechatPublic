package flyweight

import "fmt"

const (
	redTeaType   = "redTea"
	greenTeaType = "greenTea"
)

var (
	teaFactorySingleInstance = &teaFactory{
		teaMap: make(map[string]tea),
	}
)

type teaFactory struct {
	teaMap map[string]tea
}

func (d *teaFactory) getTeaByType(teaType string) (tea, error) {
	if d.teaMap[teaType] != nil {
		return d.teaMap[teaType], nil
	}

	if teaType == redTeaType {
		d.teaMap[teaType] = newRedTea()
		return d.teaMap[teaType], nil
	}

	if teaType == greenTeaType {
		d.teaMap[teaType] = newGreenTea()
		return d.teaMap[teaType], nil
	}

	return nil, fmt.Errorf("wrong tea type passed")
}

func getTeaFactorySingleInstance() *teaFactory {
	return teaFactorySingleInstance
}
