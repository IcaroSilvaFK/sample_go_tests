package utils

import "testing"

func TestReverse(t *testing.T) {

	str := Reverse("icaro")

	if str != "oraci" {
		t.Errorf("Expect rocai but got recived %s", str)
	}

}
