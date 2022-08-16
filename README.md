# QuAC-Assembler
Assembles QuAC assembly into 16-bit QuAC v1.0 ISA binaries

## Build

```shell
go build src/assembler.go
```

## Usage

You'll need to build the assembler first before doing this, see the above section for details on that.

```shell
./assembler <source assembly> <destination binary>
```

### Example

```shell
./assembler test.S result.bin
```

This will compile the following assembly

```asm
movl r1, 0x9 ; 0x0109
; Test
movl r2, 1
ldr r3, [r1]
test:
add r3, r3, r2
str r3, [r1]
jpeq test
.word 0
.word 0
.word 0
.word 0
.word 0xFF
```

Into the following binary data

```
┌────────┬─────────────────────────┬─────────────────────────┬────────┬────────┐
│00000000│ 01 09 02 01 53 10 83 32 ┊ 43 10 0f 30 00 00 00 00 │•_••S•×2┊C••00000│
│00000010│ 00 00 00 00 00 ff       ┊                         │00000×  ┊        │
└────────┴─────────────────────────┴─────────────────────────┴────────┴────────┘
```
