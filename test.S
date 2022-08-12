movl r1, 0x9 ; 0x0109
; Test
movl r2, 1
ldr r3, [r1]
test:
add r3, r3, r2
str r3, [r1]
jp eq test
.word 0
.word 0
.word 0
.word 0
.word 0xFF