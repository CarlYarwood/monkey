package object

import "testing"

func TestStringHashKey(t *testing.T) {
    hello1 := &String{Value: "Hello World"}
    hello2 := &String{Value: "Hello World"}
    diff1 := &String{Value: "My name is Carl"}
    diff2 := &String{Value: "My name is Carl"}

    if hello1.HashKey() != hello2.HashKey() {
        t.Errorf("strings with same content have different hash keys")
    }

    if diff1.HashKey() != diff2.HashKey() {
        t.Errorf("stirng with same content have differen has keys")
    }

    if hello1.HashKey() == diff1.HashKey() { 
        t.Errorf("string with different content have same hash keys")
    }
}

func TestIntegerHashKey(t *testing.T) {
    hello1 := &Integer{Value: 1}
    hello2 := &Integer{Value: 1}
    diff1 := &Integer{Value: 2}
    diff2 := &Integer{Value: 2}

    if hello1.HashKey() != hello2.HashKey() {
        t.Errorf("integers with same content have different hash keys")
    }

    if diff1.HashKey() != diff2.HashKey() {
        t.Errorf("integers with same content have differen has keys")
    }

    if hello1.HashKey() == diff1.HashKey() { 
        t.Errorf("integers with different content have same hash keys")
    }
}

func TestBooleanHashKey(t *testing.T) {
    hello1 := &Boolean{Value: true}
    hello2 := &Boolean{Value: true}
    diff1 := &Boolean{Value: false}
    diff2 := &Boolean{Value: false}

    if hello1.HashKey() != hello2.HashKey() {
        t.Errorf("boolean with same content have different hash keys")
    }

    if diff1.HashKey() != diff2.HashKey() {
        t.Errorf("boolean with same content have differen has keys")
    }

    if hello1.HashKey() == diff1.HashKey() { 
        t.Errorf("boolean with different content have same hash keys")
    }
}

