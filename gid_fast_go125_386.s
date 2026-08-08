//go:build 386 && !llgo && go1.25

#include "textflag.h"

#define get_tls(r) MOVL TLS, r

TEXT ·fastGet(SB), NOSPLIT, $0-8
	get_tls(CX)
	MOVL	0(CX)(TLS*1), AX
	MOVL	80(AX), BX
	MOVL	84(AX), CX
	MOVL	BX, ret+0(FP)
	MOVL	CX, ret+4(FP)
	RET
