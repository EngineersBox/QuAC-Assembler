# QuAC-Assembler
Assembles QuAC assembly into 16-bit QuAC v1.0 ISA binaries

The specification for QuAC v1.0 can be found here: <https://comp.anu.edu.au/courses/comp3710-uarch/resources/08-QuAC-ISA>

## Installation

The standard installation can be done with `install.sh` which will build the assembler into a binary
and store it in `/usr/bin/quac_assembler`.

```shell
./install.sh
```

If you want to build from source, simply run the following.

```shell
go build src/assembler.go
```

## Usage

You'll need to build the assembler first before doing this, see the above section for details on that.

```shell
quac_assembler <source assembly> <destination binary>
```

### Example

```shell
quac_assembler test.S result.bin
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
nop
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
