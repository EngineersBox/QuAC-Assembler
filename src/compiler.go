package main

import (
	"fmt"
	"os"

	"github.com/EngineersBox/QuAC-Compiler/antlr4"
	"github.com/EngineersBox/QuAC-Compiler/src/insn"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	var args []string = os.Args[1:]
	if len(args) != 2 {
		panic("Usage: compiler <source assembly> <destination binary>")
	}

	input, _ := antlr.NewFileStream(args[0])
	lexer := antlr4.NewQuACLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := antlr4.NewQuACParser(stream)
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	parser.BuildParseTrees = true
	tree := parser.Parse()

	var visitor insn.InsnVisitor = insn.InsnVisitor{}
	fmt.Println("Calling visitor")
	var result []uint16 = visitor.Visit(tree).([]uint16)
	var bytesResult []byte
	for _, asmCommand := range result {
		bytesResult = append(bytesResult, byte((asmCommand&0xFF00)>>8), byte(asmCommand&0x00FF))
	}

	outFile, err := os.Create(args[1])
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	_, err = outFile.Write(bytesResult)
	if err != nil {
		panic(err)
	}
}
