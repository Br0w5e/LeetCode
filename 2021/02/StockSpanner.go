package main

//901. 股票价格跨度

//单调栈
type StockSpanner struct {
	prices []int
	spans []int
}


func Constructor() StockSpanner {
	prices := make([]int, 0)
	spans := make([]int, 0)
	return StockSpanner {
		prices,
		spans,
	}
}


func (this *StockSpanner) Next(price int) int {
	s := 1
	//最大连续日数才是跨度
	for len(this.prices) > 0 && this.prices[len(this.prices)-1] <= price {
		//出栈
		this.prices = this.prices[:len(this.prices)-1]
		//加上该数字的跨度
		s += this.spans[len(this.spans)-1]
		//出栈
		this.spans = this.spans[:len(this.spans)-1]
	}
	//当前数字入栈
	this.prices = append(this.prices, price)
	//当前跨度入栈
	this.spans = append(this.spans, s)
	return s
}


/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
