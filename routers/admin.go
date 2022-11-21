package routers

import (
	"SideProject/store/models"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

// GET /
// AdminPage : admin home page
func AdminPage(w http.ResponseWriter, r *http.Request) {
	// admin: show the admin pages
	tpl := parseTemplateFile("admin/admin", "admin/admin-nav")

	// products
	
	tpl.ExecuteTemplate(w, "admin", nil)
}

// POST
// AdminAddProduct : add the new product 
func AdminAddProduct(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	// Get the upload image 
	file, handler, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	
	// create new file
	path := "./public/products/" + handler.Filename
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)

	// create image struct
	image := models.Image{
		ImageName: handler.Filename,
		ImagePath: path,
	}

	title := r.PostFormValue("title")
	price, _ := strconv.ParseFloat(r.PostFormValue("price"), 64)
	stock, _ := strconv.Atoi(r.PostFormValue("stock"))
	sale, _ := strconv.Atoi(r.PostFormValue("sale"))
	
	// create product struct
	product := models.Product{
		Title: title,
		Prices: price,
		Stocks: stock,
		Sales: sale,
		Images: []*models.Image{&image},
	}
	err = product.Create()
	if err != nil {
		log.Println(err)
	}

	// redirect to home page
	http.Redirect(w, r, "/admin", http.StatusFound)
}

