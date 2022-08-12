import re, sys

VISITOR_METHOD_REGEX = re.compile(r"(Visit\w+)\(ctx \*(\w+)\)")
TARGET_METHOD_REFEX = re.compile(r"func \(v \*InsnVisitor\) Visit\(tree antlr\.ParseTree\) interface\{\} \{[.\n]*\}")

VISIT_METHOD_PREFIX = """func (v *InsnVisitor) Visit(tree antlr.ParseTree) interface{} {
    switch val := tree.(type) {
"""
VISIT_METHOD_POSTFIX = """default:
        panic("Unknown context")
    }
}
"""
CASE_STATEMENT_TEMPLATE = """    case *antlr4.{0}:
        return v.{1}(val)
"""

def main(visitor: str) -> str:
    matches = VISITOR_METHOD_REGEX.findall(visitor)
    caseStatements = []
    for match in matches:
        caseStatements.append(CASE_STATEMENT_TEMPLATE.format(match[1], match[0]))
    return "{}{}{}".format(
        VISIT_METHOD_PREFIX,
        "\n".join(caseStatements),
        VISIT_METHOD_POSTFIX
    )

if __name__ == "__main__":
    args = sys.argv
    if (len(args) != 3):
        print("Usage: generate_visitor <go visitor interface> <target visitor implementation>")
        exit(1)
    visitorFile = open(args[1], mode="r")
    parseMethod = main(visitorFile.read())
    visitorFile.close()

    targetFile = open(args[2], mode="r")
    target: str = targetFile.read()
    targetFile.close()

    targetFile = open(args[2], mode="w")
    targetFile.write(TARGET_METHOD_REFEX.sub(parseMethod, target))
    targetFile.close()
