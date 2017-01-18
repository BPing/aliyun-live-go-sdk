package util

import "testing"

func TestCreateRandomString(t *testing.T) {
	for i := 0; i < 10; i++ {
		s := CreateRandomString()
		t.Logf("Generated Random String: %s", s)
	}
}
