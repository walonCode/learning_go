package mathutils

import "testing"

func TestMax(t *testing.T){
	xs := []float64{1,2,3,4,5,7,10}
	want := 10.0
	got := Max(xs)
	if want != got{
		t.Errorf("want %f got %f", want, got)
	}
}

func TestMin(t *testing.T){
	xs:= []float64{34,56,7,78,89,99,1,}
	want := 1.0
	got:= Min(xs)

	if want != got{
		t.Errorf("want %f got %f", want, got)
	}
}