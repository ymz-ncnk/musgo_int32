# MusGo_int32
Provides serialization of the Golang's `int32` type.

# How to use
Simply cast `int32` to `musgo_int32.Int32`:
```go
	var n int32 = 5
	// Marshal
	buf := make([]byte, musgo_int32.Int32(n).SizeMUS())
	musgo_int32.Int32(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_int32.Int32)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

