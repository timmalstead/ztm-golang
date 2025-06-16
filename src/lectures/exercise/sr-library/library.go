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
	belongsToLibrary         *Library
	booksCurrentlyCheckedOut map[string]*Book
}

type Book struct {
	bookName           string
	belongsToLibrary   *Library
	checkedOutTo       *Member
	timeLastReturned   int64
	timeLastCheckedOut int64
}

var p = fmt.Println

func unixTime() int64 {
	return time.Now().UnixNano()
}

func createLibrary(libraryName string) Library {
	return Library{
		libraryName:       libraryName,
		membershipRecords: map[string]*Member{},
		bookRecords:       map[string]*Book{},
	}
}

func (library *Library) enrollMember(memberName string) Member {
	var newMember = Member{
		memberName:               memberName,
		belongsToLibrary:         library,
		booksCurrentlyCheckedOut: map[string]*Book{},
	}

	library.membershipRecords[memberName] = &newMember

	return newMember
}

func (library *Library) donateBook(bookName string) Book {
	var newBook = Book{
		bookName:         bookName,
		belongsToLibrary: library,
	}

	library.bookRecords[bookName] = &newBook

	return newBook
}

func (book *Book) checkout(libraryMember *Member) {
	if book.checkedOutTo == nil {
		book.checkedOutTo = libraryMember
		book.timeLastCheckedOut = unixTime()
		libraryMember.booksCurrentlyCheckedOut[book.bookName] = book
	} else {
		p(book.bookName, "is already checked out by", book.checkedOutTo.memberName)
		p(book.checkedOutTo.memberName, "will need to return it before", libraryMember.memberName, "can check it out")
	}
}

func (book *Book) returnBook() {
	delete(book.checkedOutTo.booksCurrentlyCheckedOut, book.bookName)
	book.checkedOutTo = nil
	book.timeLastReturned = unixTime()
}

func (library *Library) displayLibraryRecords() {
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
	var d = centralLibrary.displayLibraryRecords

	d()

	var (
		tim    = centralLibrary.enrollMember("Tim")
		sophia = centralLibrary.enrollMember("Sophia")
		sam    = centralLibrary.enrollMember("Sam")
		anna   = centralLibrary.enrollMember("Anna")
	)

	d()

	var (
		miceAndMen = centralLibrary.donateBook("Of Mice and Men")
		poppy      = centralLibrary.donateBook("Sloppy Poppy")
		donQui     = centralLibrary.donateBook("Don Quixote")
		rebecca    = centralLibrary.donateBook("Rebecca")
	)

	d()

	poppy.checkout(&tim)
	miceAndMen.checkout(&sophia)
	donQui.checkout(&sam)
	donQui.checkout(&tim)
	rebecca.checkout(&anna)

	d()

	miceAndMen.returnBook()
	donQui.returnBook()
	rebecca.returnBook()

	d()

	miceAndMen.checkout(&tim)
	donQui.checkout(&tim)
	rebecca.checkout(&tim)

	d()

	poppy.returnBook()

	d()
}
