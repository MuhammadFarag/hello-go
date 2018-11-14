package main

import (
	"fmt"
)

func main() {

	books := make(chan *Book)

	go func() {
		defer close(books)
		for i := 0; i <= 1; i++ {
			books <- &Book{fmt.Sprintf("Author-%d", i), fmt.Sprintf("Content-%d", i)}
		}
	}()

	saveBook(printableChanToBookChan(print(bookChanToPrintableChan(books))))

}

func saveBook(books chan *Book) {
	for book := range books {
		println("Saving book for author:", book.Author)
	}
}

func print(input chan Printable) chan Printable {
	output := make(chan Printable)
	go func() {
		defer close(output)
		for printable := range input {
			println("Printing:", printable.Print())
			output <- printable
		}
	}()
	return output
}

type Book struct {
	Author  string
	Content string
}

func (b *Book) Print() string {
	return b.Content
}

type Printable interface {
	Print() string
}

func bookChanToPrintableChan(books chan *Book) chan Printable {
	printables := make(chan Printable)
	go func() {
		defer close(printables)
		for printable := range books {
			printables <- printable
		}
	}()
	return printables
}

func printableChanToBookChan(printables chan Printable) chan *Book {
	books := make(chan *Book)
	go func() {
		defer close(books)
		for printable := range printables {
			books <- printable.(*Book)
		}
	}()
	return books
}
