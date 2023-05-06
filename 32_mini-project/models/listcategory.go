package models

type ListCategory struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

type ListCategoryRepository interface {
    GetAll() ([]ListCategory, error)
}

type listCategoryRepository struct {
    dbConnection string // konfigurasi koneksi database
}

func NewListCategoryRepository(dbConnection string) ListCategoryRepository {
    return &listCategoryRepository{dbConnection}
}

func (r *listCategoryRepository) GetAll() ([]ListCategory, error) {
    // implementasi kode untuk mengambil data category dari database menggunakan dbConnection
    // contoh implementasi:
    // categories := []ListCategory{
    //     {ID: 1, Name: "Category 1"},
    //     {ID: 2, Name: "Category 2"},
    //     {ID: 3, Name: "Category 3"},
    // }
    // return categories, nil

    // karena hanya contoh, saya kembalikan data hard-coded
    categories := []ListCategory{
        {ID: 1, Name: "Nasi"},
        {ID: 2, Name: "Mie"},
        {ID: 3, Name: "Minuman"},
    }
    return categories, nil
}
