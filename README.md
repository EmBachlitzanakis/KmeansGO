# KMeans Clustering in Go

This repository contains an implementation of the KMeans clustering algorithm in Go. The program generates random 2D points, clusters them into a specified number of groups, and finds optimal centroids by minimizing the Sum of Squared Errors (SSE).

The Euclidean distance between points is calculated using a function written in Go assembly. This implementation theoretically optimizes performance by directly using low-level assembly instructions, which reduces the overhead associated with function calls and general-purpose mathematical operations.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Usage](#usage)
  - [Running the Program](#running-the-program)
  - [Example Output](#example-output)
- [KMeans Algorithm](#kmeans-algorithm)
- [Assembly Implementation](#assembly-implementation)

## Overview

The KMeans algorithm is used to partition `n` data points into `k` clusters, where each point belongs to the cluster with the nearest centroid. The algorithm works iteratively by:
1. Initializing `k` random centroids.
2. Assigning each point to the closest centroid.
3. Recomputing the centroids based on the current assignments.
4. Repeating until the centroids stabilize (or a maximum number of iterations is reached).

In this implementation, the algorithm runs multiple times, and the best centroids are selected based on the minimum SSE.

## Features

- Generates random 2D points in different clusters.
- Implements the KMeans clustering algorithm.
- Finds the best centroids based on the minimum Sum of Squared Errors (SSE).
- Saves the generated points and the best centroids to text files.

## Requirements

- Go 1.13+ installed on your system.

## Usage

### Running the Program

1. Clone this repository:

2. Build and run the program:
    ```bash
    go run main.go
    ```

3. The program will:
    - Generate random 2D points in clusters.
    - Apply the KMeans algorithm to assign these points to clusters.
    - Output the centroids of the clusters to a text file.

### Example Output

The program generates two output files:

- `Points.txt`: Contains randomly generated points in 2D space.
- `centroids.txt`: Contains the centroids of the best clusters found during the execution.

Example format of `centroids.txt`:
1.230987 1.456732 C1 0.987654 0.234567 C2 ...

## KMeans Algorithm

The KMeans algorithm is a widely used clustering technique, particularly in data mining and machine learning. This implementation follows these steps:

1. **Initialization**: Randomly choose `k` centroids within a 2D plane.
2. **Assignment**: For each data point, calculate its Euclidean distance to each centroid and assign it to the nearest one.
3. **Update**: After assignment, compute new centroids as the mean of all points in a cluster.
4. **Repeat**: Continue reassigning points and updating centroids until the centroids no longer move significantly or a maximum number of iterations is reached.

In this implementation, the algorithm is repeated 15 times, and the centroids with the lowest SSE (Sum of Squared Errors) are selected.

### Sum of Squared Errors (SSE)

The Sum of Squared Errors (SSE) is used to evaluate the quality of clusters. It measures how close the points in a cluster are to their centroid. A lower SSE means a better clustering result.

![image](https://github.com/user-attachments/assets/c3a49dcf-0a11-44d6-82d0-76bb5b3a5086)

# Euclidean Distance Implementation

## Assembly Implementation


```assembly
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
```

The assembly version of the Euclidean distance calculation manually handles each step:

- **Loading Values**: Loads the values into registers (X0, X1, X2, X3).
- **Mathematical Operations**: Performs subtraction, squares the differences, adds them, and finally computes the square root.
- **SSE Instructions**: Utilizes SSE (Streaming SIMD Extensions) instructions (`MOVSD`, `SUBSD`, `MULSD`, `ADDSD`, `SQRTSD`), which are hardware-level floating-point operations that are highly optimized on modern processors.

## Comparison: Why Assembly is Likely Faster

1. **Function Overhead**: The Go version includes calls to `math.Pow` and `math.Sqrt`, which adds overhead due to function calls. In contrast, the assembly version uses low-level instructions directly.
   
2. **Efficiency**: In the assembly code, each mathematical operation is performed using specialized CPU instructions (`MULSD`, `ADDSD`, `SQRTSD`), which are extremely fast. The Go implementation relies on general-purpose math functions that aren't as optimized for the specific case of squaring and square rooting.

3. **Avoiding Unnecessary Generality**: The Go version uses `math.Pow` to square the differences, but the assembly version directly multiplies the values, avoiding the overhead associated with the more general `math.Pow` function.

## Benchmarking to Confirm

While we can theoretically assert that the assembly version is more efficient, the only way to definitively prove which one is faster is to benchmark both implementations.
There is the calculations_test.go file inside the Calculations folder

cd Calculations
go test -bench=.

so you can see the results by yourself.

