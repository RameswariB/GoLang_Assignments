/*
Trends in technology
Group: Rameswari Bhoi,
Jagrut Gosai,
Albelis Beccea,
Gabriel Paz.
Date: 2022-01-12


2 –	In the old system of calculating sales tax in Québec, the taxes on a product were 7% for the GST, and 7.5% for the QST
 (applied after calculating the GST). Make a program that reads the unit price of a product and the quantity purchased, and 
 that displays the amounts for the GST, the QST, and the total price after taxes. 


*/

package main

import "fmt"

type Product struct{
	name string
	unitprice float64
	quantity int
	GST,QST float64
}
type taxesValue struct{
	GST,QST float64
}
func main() {

	var quantity int
	var name string
	var  price float64

	fmt.Print("Enter Product Name :")
	fmt.Scan(&name)
	fmt.Print("Enter Product price :")
	fmt.Scan(&price)
	fmt.Print("Enter Quantity :")
	fmt.Scan(&quantity)
	
	taxesValue tax = taxesValue{0.07,0.075}
	name := Product{name,price,quantity,0,0}
	name.
	fmt.Printf("Total price of the product is %.2f", totalPrice)
}
func (oneProduct *Product)calculateTaxes(taxes taxesValue){
	oneProduct.GST = oneProduct.unitprice * taxes.GST
	oneProduct.QST = (oneProduct.unitprice + oneProduct.GST) * taxes.QST
}
func 
