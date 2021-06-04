package cpu

// Represents the RAM memory available to the 6502 CPU.  Stores 8-bit
// values using a 16-bit address for a total of 65,536 possible 8-bit
// values.
type Memory interface {
	Reset()                                             // Sets all memory locations to zero
	Fetch(address uint16) (value uint8)                 // Returns the value stored at the given memory address
	Store(address uint16, value uint8) (oldValue uint8) // Stores the value at the given memory address
}

// Represents the 6502 CPU's memory using a static array of uint8's.
type BasicMemory struct {
	M             []uint8
	DisableReads  bool
	DisableWrites bool
}

// Returns a pointer to a new BasicMemory with all memory initialized
// to zero.

func NewBasicMemory(size uint32) *BasicMemory {
	return &BasicMemory{
		M: make([]uint8, size),
	}
}

// Resets all memory locations to zero
func (mem *BasicMemory) Reset() {
	for i := range mem.M {
		mem.M[i] = 0xff
	}
}

// Returns the value stored at the given memory address
func (mem *BasicMemory) Fetch(address uint16) (value uint8) {
	if mem.DisableReads {
		value = 0xff
	} else {
		value = mem.M[address]
	}

	return
}

// Stores the value at the given memory address
func (mem *BasicMemory) Store(address uint16, value uint8) (oldValue uint8) {
	if !mem.DisableWrites {
		oldValue = mem.M[address]
		mem.M[address] = value
	}

	return
}
