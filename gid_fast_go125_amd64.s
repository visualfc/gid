//go:build amd64 && !llgo && go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVQ	g, AX
	MOVQ	152(AX), AX
	MOVQ	AX, ret+0(FP)
	RET
