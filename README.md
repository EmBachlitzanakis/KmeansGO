# KMeans Clustering in Go

This repository contains an implementation of the KMeans clustering algorithm in Go. The program generates random 2D points, clusters them into a specified number of groups, and finds optimal centroids by minimizing the Sum of Squared Errors (SSE).

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Requirements](#requirements)
- [Usage](#usage)
  - [Running the Program](#running-the-program)
  - [Example Output](#example-output)
- [KMeans Algorithm](#kmeans-algorithm)

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
