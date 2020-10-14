package decorate

import "testing"

func TestComponent1_Operate(t *testing.T) {
	c1 := &Component1{}
	c1.Operate()
}

func TestDecorate1_Operate(t *testing.T) {
	d1 := &Decorate1{}
	c1 := &Component1{}
	d1.c = c1
	d1.Operate()
}