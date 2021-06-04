package cpu

import (

)

type Registers struct {
	A  uint8  // accumulator
	PC uint8 // program counter
    MAR uint8 // Memory address register
    MDR uint8 // Memory data register
    IR uint8 // Instruction register
}

type CPU struct {
    Memory Memory
    Instructions InstructionTable
    Registers Registers
}

func newCPU() {

}
