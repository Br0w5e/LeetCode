package main

//1169. 查询无效交易


//以交易名称为键，以交易信息为值建立哈希表，并逐个处理还洗白的值
import (
	"sort"
	"strconv"
	"strings"
)

type Transaction struct {
	Name   string
	Time   int
	Amount int
	City   string
	Idx    int //交易编号
}

func invalidTransactions(transactions []string) []string {
	m := make(map[string][]Transaction)
	res := make([]string, 0)
	for i, t := range transactions {
		arr := strings.Split(t, ",")
		name, city := arr[0], arr[3]
		time, _ := strconv.Atoi(arr[1])
		amount, _ := strconv.Atoi(arr[2])
		m[name] = append(m[name], Transaction{
			Name:   name,
			Time:   time,
			Amount: amount,
			City:   city,
			Idx:    i,
		})
	}

	for _, v := range m {
		//按照交易时间排序
		sort.SliceStable(v, func(i, j int) bool {
			return v[i].Time < v[j].Time
		})

		last := 0
		invalid := make(map[int]bool)
		for i := 0; i < len(v); i++ {
			if v[i].Amount > 1000 {
				invalid[v[i].Idx] = true
			}
			//last 增加的唯一标准，就是时间大于60分钟
			if v[i].Time-v[last].Time > 60 {
				last++
			}

			//遍历从last到当前位置的时间和城市是否冲突
			//["alice,20,800,mtv","alice,50,100,mtv","alice,51,100,frankfurt"]
			for j := last; j < i; j++ {
				if v[i].Time-v[j].Time <= 60 && v[i].City != v[j].City {
					invalid[v[i].Idx] = true
					invalid[v[j].Idx] = true
				}
			}
		}

		for k, _ := range invalid {
			res = append(res, transactions[k])
		}
	}

	return res
}
