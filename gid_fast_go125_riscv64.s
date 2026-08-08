//go:build riscv64 && !llgo && go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOV	g, X5
	MOV	152(X5), X5
	MOV	X5, ret+0(FP)
	RET
