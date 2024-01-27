package mmu

import "fmt"

type MemoryWriter interface {
	Read8(address uint64) byte
	Read16(address uint64) uint16
	Read32(address uint64) uint32
	Read64(address uint64) uint64
	Write(address uint64, data any) error
}

type MMU struct {
	memory    []any
	protected []bool
}

func NewMMU(size uint64) *MMU {
	mmu := new(MMU)
	mmu.memory = make([]any, size)
	mmu.protected = make([]bool, size)
	return mmu
}

func (mmu *MMU) Protect(address uint64) {
	mmu.protected[address] = true
}

func (mmu *MMU) Read8(address uint64) uint8 {
	return mmu.memory[address].(uint8)
}

func (mmu *MMU) Read16(address uint64) uint16 {
	return mmu.memory[address].(uint16)
}

func (mmu *MMU) Read32(address uint64) uint32 {
	return mmu.memory[address].(uint32)
}

func (mmu *MMU) Read64(address uint64) uint64 {
	return mmu.memory[address].(uint64)
}

func (mmu *MMU) Write(address uint64, data any) error {
	if mmu.protected[address] {
		return fmt.Errorf("invalid write: protected")
	}
	mmu.memory[address] = data
	return nil
}
