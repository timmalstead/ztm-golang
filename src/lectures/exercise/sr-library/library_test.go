package main

import "testing"

const (
	TestLibrary = "Testing Library"
	TestMember  = "Mr. Test"
	TestBook    = "Good Book"
)

func TestLibraryCreation(t *testing.T) {
	var library = createLibrary(TestLibrary)

	if library.libraryName == "" {
		t.Fatal("library name was not set correctly", library)
	}

	if len(library.bookRecords) > 0 {
		t.Fatal("library should be created with no books", library)
	}

	if len(library.membershipRecords) > 0 {
		t.Fatal("library should be created with no members", library)
	}
}

func TestMemberCreation(t *testing.T) {
	var library = createLibrary(TestLibrary)
	var testMember = library.enrollMember(TestMember)

	if testMember.belongsToLibrary.libraryName != TestLibrary {
		t.Error("member was not correctly added to library", library, testMember)
	}

	if len(library.membershipRecords) != 1 {
		t.Error("member was not correctly added to library", library, testMember)
	}

	if library.membershipRecords[TestMember].memberName != TestMember {
		t.Error("member was not correctly added to library", library, testMember)
	}
}

func TestBookDonation(t *testing.T) {
	var library = createLibrary(TestLibrary)
	var testBook = library.donateBook(TestBook)

	if len(library.bookRecords) != 1 {
		t.Error("book was not correctly added to library", library, testBook)
	}

	if library.bookRecords[TestBook].bookName != TestBook {
		t.Error("book was not correctly added to library", library, testBook)
	}
}

func TestBookCheckout(t *testing.T) {
	var library = createLibrary(TestLibrary)
	var testMember = library.enrollMember(TestMember)
	var testBook = library.donateBook(TestBook)

	testBook.checkout(&testMember)

	if testMember.booksCurrentlyCheckedOut[TestBook] == nil {
		t.Error("book was not correctly checked out", library, testMember, testBook)
	}

	if testBook.timeLastCheckedOut == 0 {
		t.Error("book was not correctly checked out", library, testMember, testBook)
	}

	if testBook.checkedOutTo == nil {
		t.Error("book was not correctly checked out", library, testMember, testBook)
	}
}

func TestBookReturn(t *testing.T) {
	var library = createLibrary(TestLibrary)
	var testMember = library.enrollMember(TestMember)
	var testBook = library.donateBook(TestBook)

	testBook.checkout(&testMember)
	testBook.returnBook()

	if testBook.checkedOutTo != nil {
		t.Error("book was not correctly returned", library, testMember, testBook)
	}

	var _, bookIsCheckedOutToTestMember = testMember.booksCurrentlyCheckedOut[TestBook]

	if bookIsCheckedOutToTestMember {
		t.Error("book was not correctly returned", library, testMember, testBook)
	}

	if testBook.timeLastReturned == 0 {
		t.Error("book was not correctly returned", library, testMember, testBook)
	}
}
