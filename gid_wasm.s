// Copyright 2026 The gid Authors. All rights reserved.
// Use of this source code is governed by the MIT license in the LICENSE file.

#include "textflag.h"

TEXT ·getg(SB), NOSPLIT, $0-8
	MOVD	g, R8
	MOVD	R8, ret+0(FP)
	RET
