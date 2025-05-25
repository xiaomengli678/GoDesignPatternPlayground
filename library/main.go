package main

import (
	"fmt"
)

var BorrowStatus map[string]string = map[string]string{
	"Available": "Available",
	"Borrowed":  "Borrowed",
	"Expired":   "Expired",
}

type Observer interface {
	Update(status string)
	GetExpectedReturnDate() int
	GetName() string
}

type LibrarySystemWithObserver struct {
	observers []Observer
}

func (l *LibrarySystemWithObserver) RegisterObserver(observer Observer) {
	l.observers = append(l.observers, observer)
}
func (l *LibrarySystemWithObserver) removeObserver(observer Observer) {
	for i, o := range l.observers {
		if o == observer {
			l.observers = append(l.observers[:i], l.observers[i+1:]...)
			break
		}
	}
}
func (l *LibrarySystemWithObserver) NotifyObserver(currentDate int) {
	for _, o := range l.observers {
		if o.GetExpectedReturnDate() < currentDate {
			o.Update(BorrowStatus["Expired"])
			fmt.Println(o.GetName())
		}
	}
}

type bookObserver struct {
	number     int
	title      string
	status     string
	returnDate int
}

func (b *bookObserver) Update(status string) {
	b.status = status
}
func (b *bookObserver) GetExpectedReturnDate() int {
	return b.returnDate
}
func (b *bookObserver) GetName() string {
	return b.title
}
func createNewBook(number int, title string, returnDate int) *bookObserver {
	return &bookObserver{number: number, title: title, status: BorrowStatus["Borrowed"], returnDate: returnDate}
}

func main() {
	book_system := LibrarySystemWithObserver{}
	book1 := createNewBook(1, "harry potter", 10)
	book2 := createNewBook(2, "kill a mocking bird", 12)
	book_system.RegisterObserver(book1)
	book_system.RegisterObserver(book2)
	today_is := 11
	book_system.NotifyObserver(today_is)
}
