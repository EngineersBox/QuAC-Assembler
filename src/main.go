package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
		".word": parseWord,
	}
	registerRegex        *regexp.Regexp = regexp.MustCompile("r\\d+")
	addressRegisterRegex *regexp.Regexp = regexp.MustCompile("\\[r\\d+\\]")
	commentRegex         *regexp.Regexp = regexp.MustCompile(";.*")
)

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
	insn |= 0b0001_0000_0000_0000
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
	insn |= 0b0100_0000_0000_0000
	return insn, nil
}

func parseLdr(args []string) (uint16, error) {
	insn, err := parseRFormatMem(args)
	if err != nil {
		return 0, nil
	}
	insn |= 0b0101_0000_0000_0000
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
		insn |= uint16(reg) << (i * 4)
	}
	return insn, nil
}

func parseAdd(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= 0b1000_0000_0000_0000
	return insn, nil
}

func parseSub(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= 0b1001_0000_0000_0000
	return insn, nil
}

func parseAnd(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= 0b1010_0000_0000_0000
	return insn, nil
}

func parseOrr(args []string) (uint16, error) {
	insn, err := parseRFormatALU(args)
	if err != nil {
		return 0, err
	}
	insn |= 0b1011_0000_0000_0000
	return insn, nil
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

func removeComments(line string) string {
	return commentRegex.ReplaceAllString(line, "")
}

func main() {
	var args []string = os.Args[1:]
	if len(args) != 2 {
		panic("Usage: compiler <in file> <out file>")
	}
	file, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var compiled []byte
	var lineNumber uint64 = 0
	for scanner.Scan() {
		var toParse string = scanner.Text()
		toParse = strings.TrimSpace(removeComments(toParse))

		if len(toParse) == 0 {
			lineNumber++
			continue
		}

		insn, err := ParseInsn(toParse)
		if err != nil {
			panic(fmt.Sprintf("[Line %d] %s\n", lineNumber, err))
		}
		// fmt.Printf("%s => %016b\n", toParse, insn)
		insnBytes := make([]byte, 2)
		binary.BigEndian.PutUint16(insnBytes, insn)

		compiled = append(compiled, insnBytes...)
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	outFile, err := os.Create(args[1])
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = outFile.Write(compiled)
	if err != nil {
		panic(err)
	}
}
