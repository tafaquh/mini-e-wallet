# Ewallet

Ewallet membutuhkan sebuah REST API untuk menghandle kebutuhan transaksi. REST API dibuat menggunakan bahasa Go.

## Cara menjalankan

```
Setup ENV
$ $env:EWALLET_DB_USER="root"
$ $env:EWALLET_DB_PASSWORD="" 
$ $env:EWALLET_DB_NAME="ewallet"

Run go
$ go run main.go
```

## Diagrams
![image](https://drive.google.com/uc?export=view&id=1j2M4IuDa2QQWbT3WA2QDFoI4GJy9cBJY)
Saya menggunakan ERD Diagram dengan model notasi seperti gambar:

* Bank - entitas untuk menyimpan informasi tentang institusi bank
* Bank Balances - entitas untuk menyimpan informasi balance dari bank itu sendiri
* Bank Balance History - entitas untuk menyimpan data history balance dari bank
* User - entitas untuk menyimpan informasi tentang user
* User Balance - entitas untuk menyimpan informasi balance dari user
* User Balance History - entitas untuk menyimpan data history balance dari user
* User Banks - entitas untuk menyimpan data bank milik user
 
Entitas Bank memiliki relasi one to one terhadap Entitas Bank Balance. Bank Balance memiliki relasi one to many terhadap Entitas Bank Balance Histories.

Entitas Bank memiliki relasi one to many terhadap Entitas User Banks.
Entitas User memiliki relasi one to many terhadap Entitas User Banks.

Entitas User memiliki relasi one to one terhadap Entitas User Balance. User Balance memiliki relasi one to many terhadap Entitas User Balance Histories.

## Tech Stacks
* Database - MySQL
* Cache - Gorm Cache untuk mempercepat pemanggilan query setelahnya
* Auth - JSON Web Token

## Seeder
Seeder telah disediakan untuk table User, Bank, Bank Balance, User Balance, dan User Bank. Password user yang diseed sudah diencripsi sebelum masuk ke database. 

## End Point

* Auth Login -- /api/v1/login type: POST
```json
Body JSON {
	"email": "tafaquh@gmail.com",
	"password": "tafaquh"
}
return request {
    "data": {
        "email": "tafaquh@gmail.com",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoidGFmYXF1aEBnbWFpbC5jb20iLCJ1c2VyIjp0cnVlLCJleHAiOjE2MDc0NTI2NTUsImlhdCI6MTYwNzI3OTg1NX0.qOqOE6siGJae_6jQCQ-8RygfOkTtM14kerxNreV6vJg"
    },
    "message": "Success login!"
}
```
* Auth Logout -- /api/v1/logout type: GET
```json
return request {
    "message": "Success logout!"
}
```
* Transfer -- /api/user/{id}/balance/transfer type: POST
```json
Header Authorization : Bearer + token
Body JSON {
	"amount": 1000,
	"user_target_id": 2
}
return request {
    "message": "transfer success",
    "your_balance": 1500
}
```
* Topup -- /api/user/{id}/balance/topup type: POST
```json
Header Authorization : Bearer + token
Body JSON {
	"balance_achieve": 1000
}
return request {
    "balance": 2000,
    "message": "topup success"
}
```
* Get User Balance -- /api/user/{id}/balance/balance-history type: GET
```json
Header Authorization : Bearer + token
return request {
    "data": {
        "ID": 0,
        "CreatedAt": "2020-12-06T22:45:00.676Z",
        "UpdatedAt": "2020-12-06T23:16:15.08Z",
        "DeletedAt": null,
        "id": 1,
        "user_id": 1,
        "User": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "id": 0,
            "username": "",
            "email": "",
            "password": ""
        },
        "balance": 1500,
        "balance_achieve": -500
    },
    "message": "Success get user balance"
}
```







