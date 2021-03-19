package main

//1603. 设计停车系统
//停车系统设置
type ParkingSystem struct {
	Big    int
	Medium int
	Small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 {
		if this.Big == 0 {
			return false
		}
		this.Big--
	} else if carType == 2 {
		if this.Medium == 0 {
			return false
		}
		this.Medium--
	} else {
		if this.Small == 0 {
			return false
		}
		this.Small--
	}
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
