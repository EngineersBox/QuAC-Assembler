package insn

import (
	"fmt"
	"strconv"

	"github.com/EngineersBox/QuAC-Compiler/antlr4"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type InsnVisitor struct {
	antlr4.QuACParserVisitor
}

func (v *InsnVisitor) Visit(tree antlr.ParseTree) interface{} {
	switch val := tree.(type) {
	case *antlr4.ParseContext:
		return v.VisitParse(val)

	case *antlr4.IFormatStatementContext:
		return v.VisitIFormatStatement(val)

	case *antlr4.RMemFormatStatementContext:
		return v.VisitRMemFormatStatement(val)

	case *antlr4.RALUFormatStatementContext:
		return v.VisitRALUFormatStatement(val)

	case *antlr4.NopStatementContext:
		return v.VisitNopStatement(val)

	case *antlr4.Pseudo2ParamStatementContext:
		return v.VisitPseudo2ParamStatement(val)

	case *antlr4.JprStatementContext:
		return v.VisitJprStatement(val)

	case *antlr4.JpmStatementContext:
		return v.VisitJpmStatement(val)

	case *antlr4.JpStatementContext:
		return v.VisitJpStatement(val)

	case *antlr4.WordStatementContext:
		return v.VisitWordStatement(val)

	case *antlr4.LabelStatementContext:
		return v.VisitLabelStatement(val)

	case *antlr4.IFormatContext:
		return v.VisitIFormat(val)

	case *antlr4.RMemFormatContext:
		return v.VisitRMemFormat(val)

	case *antlr4.RALUFormatContext:
		return v.VisitRALUFormat(val)

	case *antlr4.NopContext:
		return v.VisitNop(val)

	case *antlr4.Pseudo2ParamContext:
		return v.VisitPseudo2Param(val)

	case *antlr4.JprContext:
		return v.VisitJpr(val)

	case *antlr4.JpmContext:
		return v.VisitJpm(val)

	case *antlr4.JpContext:
		return v.VisitJp(val)

	case *antlr4.RegisterContext:
		return v.VisitRegister(val)
	default:
		panic("Unknown context")
	}
}

func (v *InsnVisitor) VisitParse(ctx *antlr4.ParseContext) interface{} {
	var result []uint16 = make([]uint16, 0)
	for _, statement := range ctx.AllStatement() {
		val := v.Visit(statement)
		fmt.Println(val)
		result = append(result, val.([]uint16)...)
	}
	return result
}

func (v *InsnVisitor) VisitIFormatStatement(ctx *antlr4.IFormatStatementContext) interface{} {
	return v.Visit(ctx.IFormat())
}

func (v *InsnVisitor) VisitRMemFormatStatement(ctx *antlr4.RMemFormatStatementContext) interface{} {
	return v.Visit(ctx.RMemFormat())
}

func (v *InsnVisitor) VisitRALUFormatStatement(ctx *antlr4.RALUFormatStatementContext) interface{} {
	return v.Visit(ctx.RALUFormat())
}

func (v *InsnVisitor) VisitNopStatement(ctx *antlr4.NopStatementContext) interface{} {
	return v.Visit(ctx.Nop())
}

func (v *InsnVisitor) VisitPseudo2ParamStatement(ctx *antlr4.Pseudo2ParamStatementContext) interface{} {
	return v.Visit(ctx.Pseudo2Param())
}

func (v *InsnVisitor) VisitJprStatement(ctx *antlr4.JprStatementContext) interface{} {
	return v.Visit(ctx.Jpr())
}

func (v *InsnVisitor) VisitJpmStatement(ctx *antlr4.JpmStatementContext) interface{} {
	return v.Visit(ctx.Jpm())
}

func (v *InsnVisitor) VisitJpStatement(ctx *antlr4.JpStatementContext) interface{} {
	return v.Visit(ctx.Jp())
}

func (v *InsnVisitor) VisitWordStatement(ctx *antlr4.WordStatementContext) interface{} {
	val, err := strconv.ParseInt(ctx.IntegerLiteral().GetText(), 0, 16)
	if err != nil {
		panic(err)
	}
	return []uint16{uint16(val)}
}

func (v *InsnVisitor) VisitLabelStatement(ctx *antlr4.LabelStatementContext) interface{} {
	return make([]uint16, 0)
}

func (v *InsnVisitor) VisitIFormat(ctx *antlr4.IFormatContext) interface{} {
	var result uint16
	if ctx.MOVL() != nil {
		result = MOVL_MASK
	} else if ctx.SETH() != nil {
		result = SETH_MASK
	} else {
		panic("Invalid I-Format instruction")
	}
	if ctx.EQ() != nil {
		result |= TRUE_CONDITION
	}
	var rd uint16 = v.Visit(ctx.Register()).(uint16)
	result |= rd << 8
	imm8, err := parseImm8(ctx.IntegerLiteral().GetText())
	if err != nil {
		panic(fmt.Errorf("Bad immediate 8-bit format: %s", err))
	}
	result |= uint16(imm8)
	fmt.Printf("I-Format: %016b\n", result)
	return []uint16{result}
}

func (v *InsnVisitor) VisitRMemFormat(ctx *antlr4.RMemFormatContext) interface{} {
	var result uint16
	if ctx.STR() != nil {
		result = STR_MASK
	} else if ctx.LDR() != nil {
		result = LDR_MASK
	} else {
		panic("Invalid R-Format memory instruction")
	}
	if ctx.EQ() != nil {
		result |= TRUE_CONDITION
	}
	var registers []antlr4.IRegisterContext = ctx.AllRegister()
	if len(registers) != 2 {
		panic(fmt.Errorf("Expected 2 registers (rd, ra), got %d", len(registers)))
	}
	result |= v.Visit(registers[0]).(uint16) << RD_REGISTER_OFFSET
	result |= v.Visit(registers[1]).(uint16) << RA_REGISTER_OFFSET
	fmt.Printf("R-Format Mem: %016b\n", result)
	return []uint16{result}
}

func (v *InsnVisitor) VisitRALUFormat(ctx *antlr4.RALUFormatContext) interface{} {
	var result uint16
	if ctx.ADD() != nil {
		result = ADD_MASK
	} else if ctx.SUB() != nil {
		result = SUB_MASK
	} else if ctx.AND() != nil {
		result = AND_MASK
	} else if ctx.ORR() != nil {
		result = ORR_MASK
	} else {
		panic("Invalid R-Format ALU instruction")
	}
	if ctx.EQ() != nil {
		result |= TRUE_CONDITION
	}
	var registers []antlr4.IRegisterContext = ctx.AllRegister()
	if len(registers) != 3 {
		panic(fmt.Errorf("Expected 3 registers (rd, ra, rb), got %d", len(registers)))
	}
	result |= v.Visit(registers[0]).(uint16) << RD_REGISTER_OFFSET
	result |= v.Visit(registers[1]).(uint16) << RA_REGISTER_OFFSET
	result |= v.Visit(registers[2]).(uint16) << RB_REGISTER_OFFSET
	fmt.Printf("R-Format ALU: %016b\n", result)
	return []uint16{result}
}

func (v *InsnVisitor) VisitNop(ctx *antlr4.NopContext) interface{} {
	return make([]uint16, 0)
}

func (v *InsnVisitor) VisitPseudo2Param(ctx *antlr4.Pseudo2ParamContext) interface{} {
	return make([]uint16, 0)
}

func (v *InsnVisitor) VisitJpr(ctx *antlr4.JprContext) interface{} {
	return make([]uint16, 0)
}

func (v *InsnVisitor) VisitJpm(ctx *antlr4.JpmContext) interface{} {
	return make([]uint16, 0)
}

func (v *InsnVisitor) VisitJp(ctx *antlr4.JpContext) interface{} {
	return make([]uint16, 0)
}

func (v *InsnVisitor) VisitRegister(ctx *antlr4.RegisterContext) interface{} {
	if ctx.RZ() != nil || ctx.R0() != nil {
		return uint16(0b000)
	} else if ctx.R1() != nil {
		return uint16(0b001)
	} else if ctx.R2() != nil {
		return uint16(0b010)
	} else if ctx.R3() != nil {
		return uint16(0b011)
	} else if ctx.R4() != nil {
		return uint16(0b100)
	} else if ctx.FL() != nil || ctx.R5() != nil {
		return uint16(0b101)
	} else if ctx.PC() != nil || ctx.R7() != nil {
		return uint16(0b111)
	}
	panic("Invalid register")
}

func parseImm8(arg string) (uint8, error) {
	imm8, err := strconv.ParseInt(arg, 0, 8)
	if err != nil {
		return 0, err
	}
	return uint8(imm8), nil
}
