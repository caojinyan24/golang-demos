package basic

import "fmt"

func InitSlice() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	slice2 := make([]*int, len(slice1))
	for index, item := range slice1 {
		slice2[index] = &item
		fmt.Println("size=", len(slice2), "index=", index)
	}
}

type MapDemoInterface interface {
	AddKey(key string, value interface{}) error
	UpdateKey(key string, value interface{}) error
	DeleteKey(key string) error
	QueryByKey(key string) (interface{}, error)
}

type MapDemo struct {
	innerMap map[string]interface{}
}

func (s MapDemo) Init() {
	fmt.Println(s.innerMap)
	s.innerMap = map[string]interface{}{"1": 11}
}
func (s MapDemo) AddKey(key string, value interface{}) error {
	fmt.Println(s.innerMap)
	s.innerMap[key] = value
	fmt.Println("addKey=", key, "value=", value)
	return nil
}
func (s MapDemo) UpdateKey(key string, value interface{}) error {
	s.innerMap[key] = value
	fmt.Println("updateKey=", key, "value=", value)
	return nil
}

func (s MapDemo) DeleteKey(key string) error {
	s.innerMap[key] = nil
	fmt.Println("deleteKey=", key)
	return nil
}

func (s MapDemo) QueryByKey(key string) (interface{}, error) {
	value, ok := s.innerMap[key]
	fmt.Println("queryKey=", key, "ok=", ok)
	return value, nil
}
func GetSmile(s MapDemoInterface) interface{} {
	s.AddKey("smile", "^_^")
	data, _ := s.QueryByKey("smile")
	return data

}

