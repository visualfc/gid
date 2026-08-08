//go:build loong64 && !llgo && go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVV	g, R4
	MOVV	152(R4), R4
	MOVV	R4, ret+0(FP)
	RET
