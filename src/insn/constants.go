package insn

const (
	MOVL_MASK = 0b0000 << 12
	SETH_MASK = 0b0001 << 12
	STR_MASK  = 0b0100 << 12
	LDR_MASK  = 0b0101 << 12
	ADD_MASK  = 0b1000 << 12
	SUB_MASK  = 0b1001 << 12
	AND_MASK  = 0b1010 << 12
	ORR_MASK  = 0b1011 << 12
)
