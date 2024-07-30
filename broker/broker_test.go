package broker

import (
	"fmt"
	"testing"
)

type Brack struct {
	b Broker
}

func TestBroker(t *testing.T) {
	b := Brack{
		b: NewKafkaBroker([]string{"a"}, nil, "aldk", nil),
	}
	fmt.Println(b)
}
