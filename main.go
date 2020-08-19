package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func connect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := "refactory"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		panic(err.Error())
	}

	return db
}

type User struct {
	ID       int    `json:"id"`
	Fullname string `json:"full_name"`
	Created  string `json:"created_at"`
	Updated  string `json:"updated_at"`
}

type Product struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Variant string `json:"variant"`
	Price   int    `json:"price"`
	Status  int    `json:"status"`
	Created string `json:"created_at"`
	Updated string `json:"updated_at"`
}

type Cart struct {
	ID      int `json:"id"`
	User    int `json:"user_id"`
	Product int `json:"product_id"`
}

type ResponseUser struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User
}

type ResponseUsers struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User
}

type ResponseProduct struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Product
}

type ResponseProducts struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Product
}

type ResponseCarts struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Cart
}

var users []User
var products []Product
var carts []Cart

// fungsi membuat user
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var response ResponseUser
	var ID int64

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	fullName := r.FormValue("full_name")

	stmt, err := db.Prepare("INSERT INTO users (full_name) values (?)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(fullName)
	if err != nil {
		log.Print(err)
	}

	ID, err = res.LastInsertId()
	if err != nil {
		log.Print(err)
	}

	query := "select id, full_name, created_at, updated_at from users WHERE id=?"
	row := db.QueryRow(query, ID)
	err = row.Scan(&user.ID, &user.Fullname, &user.Created, &user.Updated)
	if err != nil {
		log.Fatal(err.Error())
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// fungsi membuat produk
func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	var response ResponseProduct
	var ID int64

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	variant := r.FormValue("variant")
	price := r.FormValue("price")
	status := r.FormValue("status")

	stmt, err := db.Prepare("INSERT INTO products (name, variant, price, status) values (?, ?, ?, ?)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, variant, price, status)
	if err != nil {
		log.Print(err)
	}

	ID, err = res.LastInsertId()
	if err != nil {
		log.Print(err)
	}

	query := "select id, name, variant, price, status, created_at, updated_at from products WHERE id=?"
	row := db.QueryRow(query, ID)
	err = row.Scan(&product.ID, &product.Name, &product.Variant, &product.Price, &product.Status, &product.Created, &product.Updated)
	if err != nil {
		log.Fatal(err.Error())
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = product

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// fungsi membuat cart
func createCart(w http.ResponseWriter, r *http.Request) {
	var cart Cart
	var responseCarts ResponseCarts
	var ID int64

	params := mux.Vars(r)

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	productID := r.FormValue("product_id")

	stmt, err := db.Prepare("INSERT INTO carts (user_id, product_id) values (?, ?)")
	if err != nil {
		log.Print(err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(params["id"], productID)
	if err != nil {
		log.Print(err)
	}

	ID, err = res.LastInsertId()
	if err != nil {
		log.Print(err)
	}

	query := "select id, user_id, product_id from carts WHERE id=?"
	row := db.QueryRow(query, ID)
	err = row.Scan(&cart.ID, &cart.User, &cart.Product)
	if err != nil {
		log.Fatal(err.Error())
	}

	responseCarts.Status = 1
	responseCarts.Message = "Success"
	responseCarts.Data = cart

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseCarts)
}

// fungsi mengambil semua user
func getUsers(w http.ResponseWriter, r *http.Request) {
	var users User
	var arrUser []User
	var response ResponseUsers

	db := connect()
	defer db.Close()

	rows, err := db.Query("select id, full_name, created_at, updated_at from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.ID, &users.Fullname, &users.Created, &users.Updated); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// fungsi mengambil spesifik user
func getUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var response ResponseUser

	params := mux.Vars(r)

	db := connect()
	defer db.Close()

	query := "select id, full_name, created_at, updated_at from users WHERE id=?"
	row := db.QueryRow(query, params["id"])
	err := row.Scan(&user.ID, &user.Fullname, &user.Created, &user.Updated)
	if err != nil {
		log.Fatal(err.Error())
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// fungsi mengambil semua produk
func getProducts(w http.ResponseWriter, r *http.Request) {
	var products Product
	var arrProduct []Product
	var response ResponseProducts

	db := connect()
	defer db.Close()

	rows, err := db.Query("select id, name, variant, price, status, created_at, updated_at from products")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&products.ID, &products.Name, &products.Variant, &products.Price, &products.Status, &products.Created, &products.Updated); err != nil {
			log.Fatal(err.Error())

		} else {
			arrProduct = append(arrProduct, products)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrProduct

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/product", createProduct).Methods("POST")
	router.HandleFunc("/user/{id}/cart", createCart).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")
	router.HandleFunc("/products", getProducts).Methods("GET")

	http.ListenAndServe(":8000", router)
}
