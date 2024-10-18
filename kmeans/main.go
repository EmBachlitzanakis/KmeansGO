package main

import (
	"Kmeans/Calculations"
	"Kmeans/CreateFile"
	"Kmeans/StructOfData"
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	Data             = make([]StructOfData.StructOfData, 1200) // The points of S2
	NumberOfClusters = 9                                       // Number of clusters
	centroid         = make([][2]float64, NumberOfClusters)    // Centroids for each cluster
)

func kMeans() float64 {
	prevCentroid := make([][2]float64, NumberOfClusters)
	SSE := 0.0
	var SumX, SumY float64
	var count int

	for i := range centroid {
		centroid[i] = Calculations.GetRandomCentroid() // Initial
	}

	iter := 0
	for {
		for i := range Data {
			minD := math.MaxFloat64
			for j := 0; j < NumberOfClusters; j++ {
				EucDistance := Calculations.EucliDistance([2]float64(Data[i].Data), centroid[j])
				if EucDistance <= minD {
					minD = EucDistance
					Data[i].Cluster = j
				}
			}
		}

		for i := 0; i < NumberOfClusters; i++ {
			for j := range Data {
				if Data[j].Cluster == i {
					SumX += Data[j].Data[0]
					SumY += Data[j].Data[1]
					count++
				}
			}
			if count != 0 {
				centroid[i][0] = SumX / float64(count)
				centroid[i][1] = SumY / float64(count)
			}
			SumX, SumY = 0, 0
			count = 0
		}

		if compareCentroids(prevCentroid, centroid) || iter > 1000 {
			break
		}

		copy(prevCentroid, centroid)
		iter++
	}

	for i := 0; i < NumberOfClusters; i++ {
		for j := range Data {
			if Data[j].Cluster == i {
				SSE += math.Pow(Calculations.EucliDistance([2]float64(Data[j].Data), centroid[i]), 2)
			}
		}
	}

	return SSE
}

func compareCentroids(prevCentroid, currCentroid [][2]float64) bool {
	for i := range prevCentroid {
		for j := 0; j < 2; j++ {
			if prevCentroid[i][j] != currCentroid[i][j] {
				return false
			}
		}
	}
	return true
}

func main() {
	err := CreateFile.CreateFile(1200, "Points.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	file, err := os.Open("Points.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		Points := strings.Split(line, " ")
		x1, _ := strconv.ParseFloat(Points[0], 64)
		x2, _ := strconv.ParseFloat(Points[1], 64)
		Data[row] = StructOfData.StructOfData{Data: []float64{x1, x2}}
		row++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	bestCentroid := make([][2]float64, NumberOfClusters)
	var sse, minSSE float64 = 0, math.MaxFloat64

	for i := 0; i < 15; i++ {
		sse = kMeans()
		if sse <= minSSE {
			minSSE = sse
			copy(bestCentroid, centroid)
		}
	}

	centroidFile, err := os.Create("centroids.txt")
	if err != nil {
		fmt.Println("Error creating centroids file:", err)
		return
	}
	defer centroidFile.Close()

	writer := bufio.NewWriter(centroidFile)
	for i := range bestCentroid {
		writer.WriteString(fmt.Sprintf("%f %f \n", bestCentroid[i][0], bestCentroid[i][1]))
	}
	writer.Flush()

	fmt.Printf("Best Sum Of Squared Errors is: %f\n", minSSE)
}
