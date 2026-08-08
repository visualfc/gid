//go:build arm && !llgo && go1.23 && !go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVW	g, R0
	MOVW	84(R0), R1
	MOVW	88(R0), R2
	MOVW	R1, ret+0(FP)
	MOVW	R2, ret+4(FP)
	RET
