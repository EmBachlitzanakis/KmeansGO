TEXT ·EucliDistance(SB), $0
    // Load input points (x and y)
    MOVSD x+0(FP), X0          // Load x[0]
    MOVSD x+8(FP), X1          // Load x[1]
    MOVSD y+16(FP), X2         // Load y[0]
    MOVSD y+24(FP), X3         // Load y[1]

    // Subtract corresponding coordinates
    SUBSD X2, X0               // x[0] - y[0]
    SUBSD X3, X1               // x[1] - y[1]

    // Square the differences
    MULSD X0, X0               // (x[0] - y[0])^2
    MULSD X1, X1               // (x[1] - y[1])^2

    // Add the squares
    ADDSD X1, X0               // (x[0] - y[0])^2 + (x[1] - y[1])^2

    // Compute the square root
    SQRTSD X0, X0              // sqrt((x[0] - y[0])^2 + (x[1] - y[1])^2)

    // Return the result
    MOVSD X0, ret+32(FP)
    RET
