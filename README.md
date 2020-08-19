# golang-test-2-refactory-id
Golang Test 2 Refactory.id  

Langkah 1 :  
Import database dari file db.sql

Langkah 2 :  
Jalankan `go build`

Langkah 3:  
Jalankan `go run main.go`

### Cara Mengakses API

1. POST `http://localhost:8000/user`  
Untuk header diberi nilai Content-Type: multipart/form-data  
Untuk body diisi field full_name  

2. POST `http://localhost:8000/product`  
Untuk header diberi nilai Content-Type: multipart/form-data  
Untuk body diisi field name, variant, price, status

3. POST `http://localhost:8000/user/{id}/cart`  
Pada link masukkan nilai id pada path {id}  
Untuk header diberi nilai Content-Type: multipart/form-data  
Untuk body diisi field product_id

4. GET `http://localhost:8000/users`

5. GET `http://localhost:8000/user/{id}`

6. GET `http://localhost:8000/products`

