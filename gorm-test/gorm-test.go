package main

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	print("This is gorm-test")
}
