package platform

import "testing"

func TestNewPlatform(t *testing.T) {
	p := NewPlatform()
	if len(p.Modules) != 1 {
		t.Fail()
	}
	t.Log(p.Modules[0].Data)
}
