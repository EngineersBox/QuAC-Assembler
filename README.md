# QuAC-Compiler
Compiles QuAC 16bit assembly into QuAC v1.0 ISA binaries

## Build

```shell
go build src/main.go
```

## Usage

You'll need to build the compiler first before doing this, see the above section for details on that.

```shell
./main <source asm> <dest binary>
```

### Example

```shell
./main test.S result.bin
```

