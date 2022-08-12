#!/bin/bash

cd antlr4
antlr4 -Dlanguage=Go -package antlr4 QuACLexer.g4 QuACParser.g4
cd -
