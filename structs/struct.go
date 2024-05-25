package structs

type Segitiga struct {
	Alas   float64 `form:"alas" json:"alas"`
	Tinggi float64 `form:"tinggi" json:"tinggi"`
	Hitung string  `form:"hitung" json:"hitung"`
}

type Persegi struct {
	Sisi   int64  `form:"sisi" json:"sisi"`
	Hitung string `form:"hitung" json:"hitung"`
}

type PersegiPanjang struct {
	Panjang int64  `form:"panjang" json:"panjang"`
	Lebar   int64  `form:"lebar" json:"lebar"`
	Hitung  string `form:"hitung" json:"hitung"`
}

type Lingkaran struct {
	Jarijari int64  `form:"jarijari" json:"jarijari"`
	Hitung   string `form:"hitung" json:"hitung"`
}

type Category struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Book struct {
	Id           int64  `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Image_url    string `json:"image_url"`
	Release_year int64  `json:"release_year"`
	Price        string `json:"price"`
	Total_page   int64  `json:"total_page"`
	Thickness    string `json:"thickness"`
	Category_id  int64  `json:"category_id"`
	Created_at   string `json:"created_at"`
	Updated_at   string `json:"updated_at"`
}
