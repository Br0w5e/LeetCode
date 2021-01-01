package main

import "math"

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ids := map[string]int{}
	//创建哈希表
	for i, word := range wordList {
		ids[word] = i
	}
	//将开始单词加入哈希表
	if _, ok := ids[beginWord]; !ok{
		wordList = append(wordList, beginWord)
		ids[beginWord] = len(wordList)-1
	}
	//判断结束单词是否在表中
	if _, ok := ids[endWord]; !ok {
		return [][]string{}
	}
	n := len(wordList)
	edges := make([][]int, len(wordList))
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			if transformCheck(wordList[i], wordList[j]) {
				edges[i] = append(edges[i], j)
				edges[j] = append(edges[j], i)
			}
		}
	}
	res := [][]string{}
	cost := make([]int, n)
	queue := [][]int{[]int{ids[beginWord]}}
	for i := 0; i < n; i++ {
		cost[i] = math.MaxInt32
	}
	cost[ids[beginWord]] = 0

	for i := 0; i < len(queue); i++ {
		now := queue[i]
		last := now[len(now)-1]
		if last == ids[endWord] {
			tmp := []string{}
			for _, index := range now {
				tmp = append(tmp, wordList[index])
			}
			res = append(res, tmp)
		} else {
			for _, to := range edges[last] {
				if cost[last]+1 <= cost[to] {
					cost[to] = cost[last]+1
					tmp := make([]int, len(now))
					copy(tmp, now)
					tmp = append(tmp, to)
					queue = append(queue, tmp)
				}
			}
		}
	}
	return res
}

func transformCheck(from string, to string) bool {
	for i := 0; i < len(from); i++ {
		if from[i] != to[i] {
			return from[i+1:] == to[i+1:]
		}
	}
	return false
}

//DFS
func findLadders(beginWord string, endWord string, wordList []string) [][]string {

	var set =make(map[string]bool)
	for _,str := range wordList{
		set[str]=false
	}
	mp := make(map[string]map[string]bool)

	if _,ok :=set[endWord];!ok{
		return nil
	}
	beginWords :=make( map[string]bool )
	endWords :=make( map[string]bool )
	beginWords[beginWord]=true
	endWords[endWord]=true

	var ret [][]string
	if !buildTree(true,beginWords,endWords,set,mp){
		return nil
	}
	DFS(mp,&ret,beginWord,endWord,[]string{beginWord})
	return ret
}

func DFS(mp map[string]map[string]bool,ret *[][]string,begins string,end string,cache []string)  {

	if sons, ok := mp[begins];ok {
		for k,_:= range sons{
			cache=append(cache,k)
			if k == end {
				var new=make([]string,len(cache))
				copy(new,cache)
				*ret=append(*ret,new)
			}
			DFS(mp,ret,k,end,cache)
			cache=cache[0:len(cache)-1]
		}
	}
}


func buildTree(is bool,beginWords map[string]bool, endWords map[string]bool,
	wordList map[string]bool,mp map[string]map[string]bool) bool{

	if len(beginWords)==0 {
		return false
	}
	ret:=false

	for str,_:=range beginWords{
		delete(wordList, str)
	}

	var next =map[string]bool{}

	for str,_ := range beginWords{
		var sons map[string]bool
		bytes := []byte(str)
		for index,_ := range bytes {
			old:=bytes[index]
			for i := 'a'; i <= 'z'; i++ {
				bytes[index]=byte(i)
				s:=string(bytes)

				//还有路能走 or 牵桥成功 的话，进行连线
				if _,ok:=wordList[s]; ok {
					//牵桥成功
					if _,ok:= endWords[s]; ok {
						ret=true
					}
					next[s]=true
					var root,add string
					if is {
						root =str
						add=s
					}else {
						root=s
						add=str
					}
					if _,ok := mp[root];  ok{
						sons=mp[root]
					}else{
						sons=make(map[string]bool)
						mp[root]=sons
					}
					sons[add]=true
				}
				bytes[index]=old
			}
		}

	}
	if ret {
		return ret
	}

	if len(next)<len(endWords) {
		return buildTree(is,next,endWords,wordList,mp)
	}else {
		return buildTree(!is,endWords,next,wordList,mp)
	}

}