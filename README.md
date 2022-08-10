# QuAC-Compiler
Compiles QuAC assembly into 16-bit QuAC v1.0 ISA binaries

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

This will compile the following assembly

```asm
movl r1, 0x9 ; 0x0109
; Test Comment
movl r2, 1
ldr r3, [r1]
add r3, r3, r2
str r3, [r1]
.word 0
.word 0
.word 0
.word 0
.word 0xFF
```

Into the following binary data

```
┌────────┬─────────────────────────┬─────────────────────────┬────────┬────────┐
│00000000│ 01 09 02 01 53 10 83 32 ┊ 43 10 00 00 00 00 00 00 │•_••S•×2┊C•000000│
│00000010│ 00 00 00 ff             ┊                         │000×    ┊        │
└────────┴─────────────────────────┴─────────────────────────┴────────┴────────┘
```
