//go:build s390x && !llgo && go1.25

#include "textflag.h"

TEXT ·fastGet(SB), NOSPLIT, $0-8
	MOVD	g, R2
	MOVD	152(R2), R2
	MOVD	R2, ret+0(FP)
	RET
