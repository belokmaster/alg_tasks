package main

type ProductOfNumbers struct {
	products []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{products: []int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.products = []int{1}
	} else {
		value := this.products[len(this.products)-1]
		this.products = append(this.products, num*value)
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k >= len(this.products) {
		return 0
	}

	return this.products[len(this.products)-1] / this.products[len(this.products)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
