//go:build (mips64 || mips64le) && !llgo && go1.23 && !go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVV	g, R4
	MOVV	160(R4), R4
	MOVV	R4, ret+0(FP)
	RET
