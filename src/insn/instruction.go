package insn

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/EngineersBox/QuAC-Compiler/antlr4"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type InsnVisitor struct {
	antlr4.QuACParserVisitor
}

type ParseHandler func([]string) (uint16, error)

var (
	insnHandler map[string]ParseHandler = map[string]ParseHandler{
		"movl":  parseMovl,
		"seth":  parseSeth,
		"str":   parseStr,
		"ldr":   parseLdr,
		"add":   parseAdd,
		"sub":   parseSub,
		"and":   parseAnd,
		"orr":   parseOrr,
		"jpr":   parseJpr,
		"cmp":   parseCmp,
		"nop":   parseNop,
		"jpm":   parseJpm,
		"jp":    parseJp,
		".word": parseWord,
	}
	registerRegex        *regexp.Regexp = regexp.MustCompile("r\\d+")
	addressRegisterRegex *regexp.Regexp = regexp.MustCompile("\\[r\\d+\\]")
)

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
		result = append(result, v.Visit(statement).([]uint16)...)
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
	return uint16(val)
}

func (v *InsnVisitor) VisitLabelStatement(ctx *antlr4.LabelStatementContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitIFormat(ctx *antlr4.IFormatContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitRMemFormat(ctx *antlr4.RMemFormatContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitRALUFormat(ctx *antlr4.RALUFormatContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitNop(ctx *antlr4.NopContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitPseudo2Param(ctx *antlr4.Pseudo2ParamContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitJpr(ctx *antlr4.JprContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitJpm(ctx *antlr4.JpmContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitJp(ctx *antlr4.JpContext) interface{} {
	return make([]byte, 0)
}

func (v *InsnVisitor) VisitRegister(ctx *antlr4.RegisterContext) interface{} {
	return make([]byte, 0)
}

func ParseInsn(insn string) (uint16, error) {
	if len(insn) < 7 {
		return 0, fmt.Errorf("Invalid insn: %s", insn)
	}
	var splitInsn []string = strings.Split(insn, " ")
	if parseFn, ok := insnHandler[splitInsn[0]]; ok {
		return parseFn(splitInsn[1:])
	}
	return 0, fmt.Errorf("Invalid insn: %s", insn)
}

func validateIFormatArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf(
			"Expected 2 arguments, got %d: \"%s\"",
			len(args),
			strings.Join(args, " "),
		)
	}
	return nil
}

func parseRegister(reg string) (uint8, error) {
	var match []byte = registerRegex.Find([]byte(reg))
	if match == nil {
		return 0, fmt.Errorf("Invalid register format")
	}
	var matchString string = strings.Trim(string(match), "r")
	intValue, err := strconv.Atoi(matchString)
	if err != nil {
		return 0, err
	} else if intValue > 7 {
		return 0, fmt.Errorf("Register does not exist: %d", intValue)
	}
	return uint8(intValue), nil
}

func validateDestRegister(rd uint8) error {
	if rd == 0 || rd > 7 {
		return fmt.Errorf("Destination register out of range: %d", rd)
	}
	return nil
}

func parseImm8(arg string) (uint8, error) {
	imm8, err := strconv.ParseInt(arg, 0, 8)
	if err != nil {
		return 0, err
	}
	return uint8(imm8), nil
}

func parseIFormat(args []string) (uint16, error) {
	if err := validateIFormatArgs(args); err != nil {
		return 0, err
	}
	rd, err := parseRegister(args[0])
	if err != nil {
		return 0, err
	} else if err = validateDestRegister(rd); err != nil {
		return 0, err
	}
	imm8, err := parseImm8(args[1])
	if err != nil {
		return 0, err
	}
	var insn uint16 = 0
	insn |= uint16(imm8)
	insn |= (uint16(rd) << 8)
	return insn, nil
}

func parseMovl(args []string) (uint16, error) {
	return parseIFormat(args)
}

func parseSeth(args []string) (uint16, error) {
	insn, err := parseIFormat(args)
	if err != nil {
		return 0, err
	}
	insn |= SETH_MASK
	return insn, nil
}

func validateRFormatMemArgs(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf(
			"Expected 2 arguments, got %d: \"%s\"",
			len(args),
			strings.Join(args, " "),
		)
	}
	return nil
}

func parseAddressRegister(reg string) (uint8, error) {
	var match []byte = addressRegisterRegex.Find([]byte(reg))
	if match == nil {
		return 0, fmt.Errorf("Invalid register format")
	}
	var matchString string = strings.Trim(
		strings.Trim(string(match), "[]"),
		"r",
	)
	intValue, err := strconv.Atoi(matchString)
	if err != nil {
		return 0, err
	} else if intValue > 7 {
		return 0, fmt.Errorf("Register does not exist: %d", intValue)
	}
	return uint8(intValue), nil
}

func parseRFormatMem(args []string) (uint16, error) {
	if err := validateRFormatMemArgs(args); err != nil {
		return 0, err
	}
	rd, err := parseRegister(args[0])
	if err != nil {
		return 0, err
	} else if err = validateDestRegister(rd); err != nil {
		return 0, err
	}
	ra, err := parseAddressRegister(args[1])
	if err != nil {
		return 0, nil
	}
	var insn uint16 = 0
	insn |= (uint16(ra) << 4)
	insn |= (uint16(rd) << 8)
	return insn, nil
}

func parseStr(args []string) (uint16, error) {
	insn, err := parseRFormatMem(args)
	if err != nil {
		return 0, nil
	}
	insn |= STR_MASK
	return insn, nil
}

func parseLdr(args []string) (uint16, error) {
	insn, err := parseRFormatMem(args)
	if err != nil {
		return 0, nil
	}
	insn |= LDR_MASK
	return insn, nil
}

func validateRFormatALUArgs(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf(
			"Expected 3 arguments, got %d: \"%s\"",
			len(args),
			strings.Join(args, " "),
		)
	}
	return nil
}

func parseRFormatALU(args []string) (uint16, error) {
	if err := validateRFormatALUArgs(args); err != nil {
		return 0, nil
	}
	var insn uint16 = 0
	var registers []uint8 = make([]uint8, 3)
	for i, arg := range args {
		reg, err := parseRegister(arg)
		if err != nil {
			return 0, err
		} else if err = validateDestRegister(reg); err != nil {
			return 0, err
		}
		registers[i] = reg
		insn |= uint16(reg) << ((len(args) - i - 1) * 4)
	}
	return insn, nil
}

func parseAdd(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= ADD_MASK
	return insn, nil
}

func parseSub(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= SUB_MASK
	return insn, nil
}

func parseAnd(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= AND_MASK
	return insn, nil
}

func parseOrr(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= ORR_MASK
	return insn, nil
}

func parseJpr(args []string) (uint16, error) {
	return 0, nil
}
func parseCmp(args []string) (uint16, error) {
	return 0, nil
}
func parseNop(args []string) (uint16, error) {
	return 0, nil
}
func parseJpm(args []string) (uint16, error) {
	return 0, nil
}
func parseJp(args []string) (uint16, error) {
	return 0, nil
}

func parseWord(args []string) (uint16, error) {
	if len(args) != 1 {
		return 0, fmt.Errorf("Expected 1 argument, got %d", len(args))
	}
	word, err := strconv.ParseInt(args[0], 0, 16)
	if err != nil {
		return 0, err
	}
	return uint16(word), err
}
