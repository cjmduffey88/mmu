package mmu

import "testing"

func TestMMU_Protect(t *testing.T) {
	mmu := NewMMU(0x100)
	mmu.Protect(0x00)
	err := mmu.Write(0x00, 0x00)
	if err == nil {
		t.Errorf("expected invalid write: protected")
	}
}
