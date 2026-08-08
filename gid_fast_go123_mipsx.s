//go:build (mips || mipsle) && !llgo && go1.23 && !go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVW	g, R4
	MOVW	84(R4), R5
	MOVW	88(R4), R6
	MOVW	R5, ret+0(FP)
	MOVW	R6, ret+4(FP)
	RET
