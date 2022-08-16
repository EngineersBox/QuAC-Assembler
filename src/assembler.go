package main

import (
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

	listener := insn.NewLabelListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	var visitor insn.InsnVisitor = insn.NewInsnVisitor(listener.Labels)
	var result []uint16 = visitor.Visit(tree).([]uint16)
	var bytesResult []byte
	for _, asmCommand := range result {
		bytesResult = append(bytesResult, byte((asmCommand&0xFF00)>>8), byte(asmCommand&0x00FF))
	}

	outFile, err := os.Create(args[1])
	if err != nil {
		panic(err)
	}
	defer func(outFile *os.File) {
		err := outFile.Close()
		if err != nil {

		}
	}(outFile)

	// In go 1.16 for some reason this re-orders bytes, but only the first 8 bytes over the 16 byte threshold. This is fixed in go1.17 and above
	_, err = outFile.Write(bytesResult)
	if err != nil {
		panic(err)
	}
}
