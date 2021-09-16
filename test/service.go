package test

import (
	"fmt"
	"github.com/caojinyan24/golang-demos/basic"
	"sort"
)

func TestParamName() {
	fmt.Println(basic.PLATFORM_APP_IDENTITY_TYPE_UID)
}

type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}
