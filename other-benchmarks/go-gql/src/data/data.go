package data

import (
	"github.com/fibs7000/go-graphql-benchmark/graph/model"
	"github.com/jaswdr/faker"
)

func GenData() []*model.Author {
	faker := faker.New()

	var authors []*model.Author
	for i := 0; i < 20; i++ {
		var books []*model.Book

		for k := 0; k < 3; k++ {
			book := model.Book{
				ID:       faker.UUID().V4(),
				Name:     faker.Person().LastName(),
				NumPages: faker.IntBetween(100, 10000),
			}
			books = append(books, &book)
		}

		author := model.Author{
			ID:      faker.UUID().V4(),
			Name:    faker.Person().FirstName(),
			Company: faker.Company().BS(),
			Books:   books,
		}
		authors = append(authors, &author)
	}

	return authors
}
