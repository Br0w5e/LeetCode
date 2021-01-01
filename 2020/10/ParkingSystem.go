package main

//1603. 设计停车系统
//停车系统设置
type ParkingSystem struct {
	big, medium, small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{big, medium, small}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big != 0 {
			this.big--
			return true
		}
	case 2:
		if this.medium != 0 {
			this.medium--
			return true
		}
	case 3:
		if this.small != 0 {
			this.small--
			return true
		}
	default:
		return false
	}
	return false
}



/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */