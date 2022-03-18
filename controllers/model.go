package controllers

import "github.com/dgrijalva/jwt-go"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address`
	UserType int    `json:"type"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type UsersResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

type Product struct {
	ID    int    `json:"id`
	Name  string `json:"name"`
	Price int    `json:"price`
}

type ProductResponse struct {
	Status  int     `json:"status"`
	Message string  `json:"message"`
	Data    Product `json:"data"`
}

type ProductsResponse struct {
	Status  int       `json:"status"`
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}

type Transaction struct {
	ID        int `json:"id"`
	UserID    int `json:"userID"`
	ProductID int `json:"productID"`
	Quantity  int `json:"qyt"`
}

type TransactionResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    Transaction `json:"data"`
}

type TransactionsResponse struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []Transaction `json:"data"`
}

type GeneralResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type DetailedTransaction struct {
	ID          int     `json:"transaction_id"`
	DataUser    User    `json:"User"`
	DataProduct Product `json:"Product"`
	Quantity    int     `json:"quantity"`
}

type DetailedTransactionResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    DetailedTransaction `json:"Data"`
}

type DetailedTransactionsResponse struct {
	Status  int                   `json:"status"`
	Message string                `json:"message"`
	Data    []DetailedTransaction `json:"Data"`
}

type Claims struct {
	ID       int    `json:"id`
	Name     string `json:"name"`
	UserType int    `json:"user_type"`
	jwt.StandardClaims
}
