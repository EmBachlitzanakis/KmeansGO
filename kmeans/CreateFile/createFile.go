package CreateFile

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// CreateFile creates a file and writes random points into it
func CreateFile(numberOfPoints int, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	rand.Seed(time.Now().UnixNano())

	var x1, x2 float64

	for i := 0; i < 150; i++ {
		x1 = 0.80 + rand.Float64()*(1.2-0.80)
		x2 = 0.80 + rand.Float64()*(1.2-0.80)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 150; i < 300; i++ {
		x1 = 0 + rand.Float64()*(0.5-0)
		x2 = 0 + rand.Float64()*(0.5-0)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 300; i < 450; i++ {
		x1 = 0 + rand.Float64()*(0.5-0)
		x2 = 1.5 + rand.Float64()*(2-1.5)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 450; i < 600; i++ {
		x1 = 1.5 + rand.Float64()*(2-1.5)
		x2 = 0 + rand.Float64()*(0.5-0)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 600; i < 750; i++ {
		x1 = 1.5 + rand.Float64()*(2-1.5)
		x2 = 1.5 + rand.Float64()*(2-1.5)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 750; i < 825; i++ {
		x1 = 0.8 + rand.Float64()*(1.2-0.8)
		x2 = 0 + rand.Float64()*(0.4-0)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 825; i < 900; i++ {
		x1 = 0.8 + rand.Float64()*(1.2-0.8)
		x2 = 1.6 + rand.Float64()*(2-1.6)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 900; i < 975; i++ {
		x1 = 0.3 + rand.Float64()*(0.7-0.3)
		x2 = 0.8 + rand.Float64()*(1.2-0.8)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 975; i < 1050; i++ {
		x1 = 1.3 + rand.Float64()*(1.7-1.3)
		x2 = 0.8 + rand.Float64()*(1.2-0.8)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	for i := 1050; i < 1200; i++ {
		x1 = 0 + rand.Float64()*(2-0)
		x2 = 0 + rand.Float64()*(2-0)
		writer.WriteString(fmt.Sprintf("%f %f\n", x1, x2))
	}

	writer.Flush()
	fmt.Println("Txt File with data is created successfully")
	return nil
}
