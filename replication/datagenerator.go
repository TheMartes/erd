package replication

import "github.com/bxcodec/faker/v3"

type FakeData []string

func GenerateFakeData(cycles int) FakeData {
	var fakeTitles FakeData

	for i := 0; i < cycles; i++ {
		fakeTitles = append(fakeTitles, faker.Word())
	}

	return fakeTitles
}
