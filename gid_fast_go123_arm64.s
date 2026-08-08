//go:build arm64 && !llgo && go1.23 && !go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVD	g, R8
	MOVD	160(R8), R8
	MOVD	R8, ret+0(FP)
	RET
