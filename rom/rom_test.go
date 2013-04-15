package rom

import (
	"testing"
)

func tAdd(in1, in2, out string, t *testing.T) {
	f := Add
	ans := f(in1, in2)
	if ans != out {
		t.Errorf("Add(%v, %v) = %v, wanted %v", in1, in2, ans, out)
	}
}

func TestConversion0(t *testing.T) {
	tAdd("VIII", "I", "IX", t)
}

func TestConversion1(t *testing.T) {
	tAdd("XVIII", "II", "XX", t)
}

func TestConversion2(t *testing.T) {
	tAdd("II", "V", "VII", t)
}

func TestConversion3(t *testing.T) {
	tAdd("XXX", "VIII", "XXXVIII", t)
}

func TestConversion4(t *testing.T) {
	tAdd("XLIV", "C", "CXLIV", t)
}

func TestConversion5(t *testing.T) {
	tAdd("XXIX", "IV", "XXXIII", t)
}

func TestConversion6(t *testing.T) {
	tAdd("M", "III", "MIII", t)
}

func TestConversion7(t *testing.T) {
	tAdd("MI", "V", "MVI", t)
}
