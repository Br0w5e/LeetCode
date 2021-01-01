package main

//731. 我的日程安排表 II

//这种题好像真的不太会，加强一下
type MyCalendarTwo struct {
	list  [][]int // [[start,end)]
	cross [][]int // [[start,end)]
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this MyCalendarTwo) GetCrossTime(s1,e1,s2,e2 int) (s,e int,ok bool) {
	s ,e = s1, e1
	if s < s2 {
		s = s2
	}
	if e > e2 {
		e = e2
	}
	return s,e,s < e
}

func (this *MyCalendarTwo) Book(start int, end int) bool {

	for i := 0; i < len(this.cross); i++ {
		_,_,ok := this.GetCrossTime(this.cross[i][0],this.cross[i][1],start,end)
		if ok {
			return false
		}
	}

	for i := 0; i < len(this.list); i++ {
		s,e,ok := this.GetCrossTime(this.list[i][0],this.list[i][1],start,end)
		if ok {
			this.cross = append(this.cross,[]int{s,e})
		}
	}

	this.list = append(this.list, []int{start, end})
	return true
}


/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
