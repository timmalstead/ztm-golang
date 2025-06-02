//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Library struct {
	libraryName       string
	membershipRecords map[string]*Member
	bookRecords       map[string]*Book
}

type Member struct {
	memberName               string
	booksCurrentlyCheckedOut map[string]*Book
}

type Book struct {
	bookName           string
	checkedOutTo       *Member
	timeLastReturned   int64
	timeLastCheckedOut int64
}

var p, unixTime = fmt.Println, time.Now().UnixNano

func createLibrary(libraryName string) Library {
	return Library{
		libraryName:       libraryName,
		membershipRecords: map[string]*Member{},
		bookRecords:       map[string]*Book{},
	}
}

func createMember(memberName string) Member {
	return Member{
		memberName:               memberName,
		booksCurrentlyCheckedOut: map[string]*Book{},
	}
}

func createBook(bookName string) Book {
	return Book{bookName: bookName}
}

func checkoutBook(book *Book, libraryMember *Member) {
	if book.checkedOutTo == nil {
		book.checkedOutTo = libraryMember
		book.timeLastCheckedOut = unixTime()
		libraryMember.booksCurrentlyCheckedOut[book.bookName] = book
	} else {
		p(book.bookName, "is already checked out by", book.checkedOutTo.memberName)
		p(book.checkedOutTo.memberName, "will need to return it before", libraryMember.memberName, "can check it out")
	}
}

func returnBook(book *Book) {
	delete(book.checkedOutTo.booksCurrentlyCheckedOut, book.bookName)
	book.checkedOutTo = nil
	book.timeLastReturned = unixTime()
}

func membershipDrive(library *Library, members []*Member) {
	for _, member := range members {
		library.membershipRecords[member.memberName] = member
	}
}

func donateBooks(library *Library, books []*Book) {
	for _, book := range books {
		library.bookRecords[book.bookName] = book
	}
}

func displayLibraryRecords(library *Library) {
	p("Library status:")
	for _, member := range library.membershipRecords {
		if len(member.booksCurrentlyCheckedOut) > 0 {
			p(member.memberName, "currently has the following books checked out: ")
		} else {
			p(member.memberName, "does not have any books checked out")
		}
		for _, book := range member.booksCurrentlyCheckedOut {
			p(book.bookName, "was checked out to", book.checkedOutTo.memberName, "at", book.timeLastCheckedOut, "before that it was returned at", book.timeLastReturned)
		}
	}
}

func main() {
	var centralLibrary = createLibrary("Central Library")
	displayLibraryRecords(&centralLibrary)

	var (
		tim    = createMember("Tim")
		sophia = createMember("Sophia")
		sam    = createMember("Sam")
		anna   = createMember("Anna")
	)

	p("Membership drive!")
	membershipDrive(&centralLibrary, []*Member{&tim, &sophia, &sam, &anna})

	displayLibraryRecords(&centralLibrary)

	var (
		miceAndMen = createBook("Of Mice and Men")
		poppy      = createBook("Sloppy Poppy")
		donQui     = createBook("Don Quixote")
		rebecca    = createBook("Rebecca")
	)

	p("Book fair!")
	donateBooks(&centralLibrary, []*Book{&miceAndMen, &poppy, &donQui, &rebecca})
	displayLibraryRecords(&centralLibrary)

	checkoutBook(&poppy, &tim)
	checkoutBook(&miceAndMen, &sophia)
	checkoutBook(&donQui, &sam)
	checkoutBook(&donQui, &tim)
	checkoutBook(&rebecca, &anna)

	displayLibraryRecords(&centralLibrary)

	returnBook(&miceAndMen)
	returnBook(&donQui)
	returnBook(&rebecca)

	displayLibraryRecords(&centralLibrary)

	checkoutBook(&miceAndMen, &tim)
	checkoutBook(&donQui, &tim)
	checkoutBook(&rebecca, &tim)

	displayLibraryRecords(&centralLibrary)

	returnBook(&poppy)
	checkoutBook(&poppy, &tim)

	displayLibraryRecords(&centralLibrary)
}
