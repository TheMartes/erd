package helpers

import (
	"runtime"
)

// Chunk of data
type Chunk struct {
	Data []string
}

// SplitIntoChunks will split data into # of cores to help reduce stress
func SplitIntoChunks(data []string) []Chunk {
	numberOfChunks := runtime.NumCPU()
	len := len(data)

	chunkLen := len / numberOfChunks
	threshold := 0

	output := make([]Chunk, numberOfChunks)

	for chunk := 0; chunk < numberOfChunks; chunk++ {
		output[chunk] = Chunk{Data: []string{}}

		for i := 0; i < chunkLen; i++ {
			output[chunk].Data = append(output[chunk].Data, data[threshold+i])
		}

		threshold = threshold + chunkLen
	}

	return output
}
