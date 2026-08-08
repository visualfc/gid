//go:build (ppc64 || ppc64le) && !llgo && go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVD	g, R3
	MOVD	152(R3), R3
	MOVD	R3, ret+0(FP)
	RET
