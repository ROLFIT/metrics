// expvar_test
package metrics

import (
	"sync"
	"testing"
)

func TestIntSet(t *testing.T) {
	v := NewInt("var_int_set", "", "", "")

	val := int64(12412331453)
	v.Set(val)

	res := v.Get()
	if res != val {
		t.Errorf("Get should be \"%v\", was \"%v\"", val, res)
	}
}

func TestIntAdd(t *testing.T) {
	v := NewInt("var_int_add", "", "", "")

	val := int64(12412331453)
	v.Add(val)

	res := v.Get()
	if res != val {
		t.Errorf("Get should be \"%v\", was \"%v\"", val, res)
	}

	val1 := int64(-254545)
	Add("var_int_add", val1)

	res = v.Get()
	if res != val+val1 {
		t.Errorf("Get should be \"%v\", was \"%v\"", val+val1, res)
	}
}

func TestFloatSet(t *testing.T) {
	v := NewFloat("var_float_set", "", "", "")

	val := float64(12412.331453)
	v.Set(val)

	res := v.Get()
	if res != val {
		t.Errorf("Get should be \"%v\", was \"%v\"", val, res)
	}
}

func TestFloatAdd(t *testing.T) {
	v := NewFloat("var_float_add", "", "", "")

	val := float64(12412.331453)
	v.Add(val)

	res := v.Get()
	if res != val {
		t.Errorf("Get should be \"%v\", was \"%v\"", val, res)
	}
}

func TestI1000(t *testing.T) {
	v := NewInt("var1", "Variable 1", "Bytes", "b")

	for j := 0; j < 1000; j++ {
		Add("var1", 1)
	}

	res := v.Get()
	if res != 1000 {
		t.Errorf("Get should be \"%v\", was \"%v\"", 1000, res)
	}
}

func TestI10x1000(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(10)
	v := NewInt("var2", "Variable 1", "Bytes", "b")

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10000; j++ {
				Add("var2", 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	res := v.Get()
	if res != 100000 {
		t.Errorf("Get should be \"%v\", was \"%v\"", 10000, res)
	}
}

func TestF1000(t *testing.T) {
	v := NewFloat("var3", "Variable 1", "Bytes", "b")

	var f float64
	delta := float64(1) / 3
	for j := 0; j < 1000; j++ {
		AddFloat("var3", delta)
		f = f + delta

	}
	res := v.Get()
	if res != f {
		t.Errorf("Get should be \"%v\", was \"%v\"", f, res)
	}
}
