package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func main() {
	// membuat channel untuk menampung data produk
	products := make(chan Product, 10)

	// menjalankan goroutine untuk mengambil data produk dari REST API
	go func() {
		resp, err := http.Get("https://fakestoreapi.com/products")
		if err != nil {
			fmt.Println(err)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		// mengkonversi data JSON menjadi slice struct Product
		var productList []Product
		err = json.Unmarshal(body, &productList)
		if err != nil {
			fmt.Println(err)
			return
		}

		// mengirim data produk ke channel
		for _, product := range productList {
			products <- product
		}

		close(products)
	}()

	// menampilkan data produk dari channel
	fmt.Println("products data \n===")
	for product := range products {
		fmt.Printf("tittle : %s\nprice : %.2f\ncategory : %s\n===\n", product.Title, product.Price, product.Category)
	}
}