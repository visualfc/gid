//go:build 386 && !llgo && go1.23 && !go1.25

#include "textflag.h"

#define get_tls(r) MOVL TLS, r
#define g_tls(r) 0(r)(TLS*1)

TEXT ·fastGet(SB), NOSPLIT, $0-8
	get_tls(CX)
	MOVL	g_tls(CX), AX
	MOVL	84(AX), BX
	MOVL	88(AX), CX
	MOVL	BX, ret+0(FP)
	MOVL	CX, ret+4(FP)
	RET
