package models

type Book struct {
    ID      int
    Title   string
    Author  string
    Publisher string
}

var Books []Book

func GetAllBooks() []Book {
    return Books
}

func GetBookById(id int) (*Book, error) {
    for _, book := range Books {
        if book.ID == id {
            return &book, nil
        }
    }
    return nil, errors.New("Book not found")
}

func CreateBook(book Book) {
    Books = append(Books, book)
}

func UpdateBook(id int, book Book) error {
    for i, b := range Books {
        if b.ID == id {
            Books[i] = book
            return nil
        }
    }
    return errors.New("Book not found")
}

func DeleteBook(id int) error {
    for i, book := range Books {
        if book.ID == id {
            Books = append(Books[:i], Books[i+1:]...)
            return nil
        }
    }
    return errors.New("Book not found")
}
