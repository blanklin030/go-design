package strategy

import "fmt"

type Context struct {
	Strategy
}

type Strategy interface {
	Do()
}

type Strategy1 struct {

}

func (s *Strategy1) Do() {
	fmt.Println("operate strategy 1")
}

type Strategy2 struct {

}

func (s *Strategy2) Do() {
	fmt.Println("operate strategy 2")
}