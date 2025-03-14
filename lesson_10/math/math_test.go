package math

import (
	"fmt"
	"testing"
)

// func TestAverage(t *testing.T){
// 	v := Average([]float64{1,2,3})
// 	if v != 2.0{
// 		t.Error("Expected 2.0, got", v)
// 	}
// }

func TestAdd(t *testing.T){
	v:=Add(1,2)
	if v != 3 {
		t.Error("Expected 3, got",v)
	}else{
		fmt.Println("Test passed")
	}
}