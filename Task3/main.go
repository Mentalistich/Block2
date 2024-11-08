package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
)

type Person struct {
	Name     string
	DaysGone int
}

type Book struct {
	title  string
	author string
	year   int
	copies int
}

func (b Book) GetDescription() string {
	s := b.title + b.author + "date of publication" + strconv.Itoa(b.year) + "number of copies" + strconv.Itoa(b.copies)
	return s

}

type Library struct {
	Books []Book
}

func (l *Library) SearchBook(title string) Book {
	for _, b := range l.Books {
		if b.title == title {
			return b
		}
	}
	fmt.Println("Not found")
	return Book{}
}

func (l *Library) AddBook(title string, author string, year int, copies int) {
	l.Books = append(l.Books, Book{title, author, year, copies})
}

func (l Library) BorrowBook(title string) error {
	for _, b := range l.Books {
		if b.title == title {
			if b.copies-1 >= 0 {
				b.copies--
				return nil
			} else {
				return errors.New("No avaliable books in library")
			}
		}
	}
	return errors.New("There is no such book in the library")
}

func (l Library) ReturnBook(title string) error {
	for _, b := range l.Books {
		if b.title == title {
			b.copies++
			return nil

		}
	}
	return errors.New("There is no such book in the library")
}

func (l Library) BooksForUsers(title string, p []Person) map[string]int {

	ILoveBooks := make(map[string]int)
	for _, b := range l.Books {
		err := l.BorrowBook(b.title)
		if err == nil {
			ILoveBooks[title]++
		}
	}
	return ILoveBooks
}

func (p *Person) CheckDebt() string {
	if p.DaysGone > 30 {
		return "Dear reader " + p.Name + " give back my book"
	} else {
		p.DaysGone++
	}

	return "All good"
}

func SortLibrary(l *Library) {
	sort.Slice(l.Books, func(i, j int) bool { return l.Books[i].copies < l.Books[j].copies })
}

func (l Library) GetTopPopularBooks(limit int) []Book {
	mostPopularBooks := make([]Book, 0)
	for _, book := range l.Books {
		if book.copies < limit {
			mostPopularBooks = append(mostPopularBooks, book)
		}
	}
	return mostPopularBooks
}

func main() {
	aleshes := []Person{}
	alesha := Person{Name: "alesha", DaysGone: 30}
	alesha2 := Person{Name: "alesha2", DaysGone: 0}
	aleshes = append(aleshes, alesha)
	aleshes = append(aleshes, alesha2)
	for i := 0; i < len(aleshes); i++ {
		fmt.Println(alesha.CheckDebt())
	}
	library := Library{}
	library.AddBook("Book 1", "Author 1", 2000, 5)
	library.AddBook("Book 2", "Author 2", 2010, 2)
	library.AddBook("Book 3", "Author 3", 2020, 7)
	fmt.Println(library.SearchBook("Book 1"))
	SortLibrary(&library)
	fmt.Println(library)
	fmt.Println(library.GetTopPopularBooks(6))
}
