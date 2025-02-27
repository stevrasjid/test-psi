# test-psi
 
Jawaban untuk test PSI menggunakan Golang + Gin.
Cara menjalankan :
1. Buka Terminal pada vscode
2. ketik "go run ." lalu enter
3. program sudah berjalan

Untuk jawaban selain nomor 3 bisa menggunakan postman:
1. URL : localhost:8080/api/number1, METHOD : POST, 
Payload : {
    "voucherDiscount" : 50,
    "productPrice" : 5000000
}

2. URL : localhost:8080/api/number2, METHOD : POST
Payload : {
    "username" : "username123"
}

3. URL : localhost:8080/api/number4?result=10&page=1, METHOD : GET, Parameter Optional
4. URL : localhost:8080/api/number5?color=maroon, Parameter Optional
