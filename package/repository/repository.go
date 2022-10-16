package repository

type Authorization struct {

}
type BooksList struct {

}
type BooksItem struct{

}
type Repository struct {
	Authorization
	BooksList
	BooksItem
}

func NewRepository()*Repository {
	return &Repository{}
}