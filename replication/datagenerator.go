package replication

import "github.com/bxcodec/faker/v3"

func GenerateFakeData(cycles int) []string {
	var fakeTitles []string

	for i := 0; i < cycles; i++ {
		fakeTitles = append(fakeTitles, faker.Word())
	}

	return fakeTitles
}
