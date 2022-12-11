package musgo_int32

import "testing"

func TestMusgoInt32(t *testing.T) {
	var n int32 = 5
	buf := make([]byte, Int32(n).SizeMUS())
	Int32(n).MarshalMUS(buf)

	var an int32
	(*Int32)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
