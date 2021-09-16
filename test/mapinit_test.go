package test

import (
	"fmt"
	"github.com/caojinyan24/golang-demos/basic"
	"github.com/caojinyan24/golang-demos/mock_basic"
	"github.com/golang/mock/gomock"

	"testing"
)

func TestMapDemo(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockMapDemo := mock_basic.NewMockMapDemo(mockCtrl)

	//mockMapDemo.EXPECT().UpdateKey(gomock.Any(), gomock.Any()).Return(nil)
	mockMapDemo.EXPECT().AddKey(gomock.Any(), gomock.Any()).Return(nil)
	mockMapDemo.EXPECT().QueryByKey("smile").Do(func(key string) {
		fmt.Println("query fake smile")
	}).Return("=_=", nil)

	data := basic.GetSmile(mockMapDemo)
	fmt.Println(data)
	if data != "=_=" {
		t.Errorf("Expected =_=, got %d", data)
	}
}
