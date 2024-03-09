package concurrency

import (
	"strconv"
)

type Product struct {
	Name, Category string
	Price          float64
}

var ProductList = []*Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

type ProductGroup []*Product

type ProductData = map[string]ProductGroup

var Products = make(ProductData)

//var Products2 = make(map[string][]*Product)

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, p := range ProductList {
		_, ok := Products[p.Category]
		if ok {
			Products[p.Category] = append(Products[p.Category], p)
			//fmt.Println(Products[p.Category])
		} else {
			//fmt.Println("ProductGroup", p)
			Products[p.Category] = ProductGroup{p}
		}
	}
}

/*
func init2() {
	for _, p := range ProductList {
		_, ok := Products2[p.Category]
		if ok {
			Products2[p.Category] = append(Products2[p.Category], p)
			fmt.Println(Products[p.Category])
		} else {
			fmt.Println("ProductGroup", p)
			Products2[p.Category] = ProductGroup{p}
		}
	}
}*/
