package basic

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestWeight(t *testing.T) {
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

	for {
		if i > 2 {
			break
		}
		i = i + 1
		instance, _ := NewWeightPicker(instances, 4).Pick()
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

type weightPicker struct {
	ins     []*Instance
	weights []int
	sum     int
}
type Instance struct {
	HostPort string
	Weight   int
	Tags     map[string]string
}

func NewWeightPicker(instances []*Instance, scale int) *weightPicker {
	weights := make([]int, len(instances))
	sum := 0
	for i, ins := range instances {
		weights[i] = ins.Weight * scale
		sum += weights[i]
	}

	return &weightPicker{
		ins:     instances,
		weights: weights,
		sum:     sum,
	}
}
func (wp *weightPicker) Pick() (*Instance, bool) {
	if wp.sum <= 0 || len(wp.ins) == 0 {
		return nil, false
	}
	rd := rand.Intn(wp.sum)
	for i := 0; i < len(wp.ins); i++ {
		if rd < wp.weights[i] {
			wp.sum -= wp.ins[i].Weight
			wp.weights[i] -= wp.ins[i].Weight
			fmt.Println("true rd=", rd, "wp=", wp.weights, "i=", i, "weight i=", wp.weights[i], "sum=", wp.sum)
			return wp.ins[i], true
		}
		rd -= wp.weights[i]
		fmt.Println("false rd=", rd, "wp=", wp.weights, "i=", i, "weight i=", wp.weights[i], "sum=", wp.sum)
	}

	return nil, false
}
