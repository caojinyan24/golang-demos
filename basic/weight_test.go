package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestWeightRandom(t *testing.T) {
	var instances = []*Instance{&Instance{
		HostPort: "1", Weight: 3,
	}, &Instance{
		HostPort: "2", Weight: 1,
	}, &Instance{
		HostPort: "3", Weight: 2,
	}}
	i := 0
	var host1count = 0
	var host2count = 0
	var host3count = 0
	for i := 0; i < len(instances); i++ {
		if i == 0 {
		} else {
			instances[i].Weight = instances[i].Weight + instances[i-1].Weight
		}
		fmt.Println(instances[i])
	}
	for {
		if i > 100 {
			break
		}
		i = i + 1
		instance, _ := NewWeightRandomPicker(instances, 4).PickRandom()
		if instance.HostPort == "1" {
			host1count = host1count + 1
		} else if instance.HostPort == "2" {
			host2count = host2count + 1
		} else {
			host3count = host3count + 1
		}
	}
	fmt.Println(host1count, host2count, host3count)
}

func NewWeightRandomPicker(instances []*Instance, scale int) *weightPicker {
	weights := make([]int, len(instances))
	sum := 0

	return &weightPicker{
		ins:     instances,
		weights: weights,
		sum:     sum,
	}
}

func (wp *weightPicker) PickRandom() (*Instance, bool) {
	rd := rand.Intn(6)
	fmt.Println(rd)
	for i := 0; i < len(wp.ins); i++ {
		fmt.Println("weight=", wp.ins[i])
		if i == 0 {
			if rd < wp.ins[i].Weight {
				return wp.ins[i], true
			}
		} else {
			if rd < wp.ins[i].Weight && rd >= wp.ins[i-1].Weight {
				return wp.ins[i], true
			}
		}
	}
	return nil, false
}
