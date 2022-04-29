package leetcode


func findSubStringMine(s string, words []string)[]int{
	// 

	var lenWord=len(words[0])
	var totalLen=lenWord*len(words)

	var cache=make(map[string]int,len(words))
	for _,w:=range words{
		cache[w]=0
	}
	
	var clean=func(c map[string]int){
		for i:=range c{
			c[i]=0
		}
	}

	func contained=func(w string)bool{
		n,ok:=cache[w]
		if !ok{
			return false
		}
		if n!=0{
			return false 
		}
		cache[w]=1
		return true
	}
	

	var wordsInString=func(i int,words []string)bool{
		if i+totalLen>len(s){
			return false
		}
		
		if len(words)==0{
			return true
		}

		if !contained(s[i:i+lenWord]){
			return false
		}


		return wordsInString(i+lenWord,)
	}



}

func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := []int{}
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}
	length, totalLen, tmpCounter := len(words[0]), len(words[0])*len(words), copyMap(counter)
	for i, start := 0, 0; i < len(s)-length+1 && start < len(s)-length+1; i++ {
		//fmt.Printf("sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
		if tmpCounter[s[i:i+length]] > 0 {
			tmpCounter[s[i:i+length]]--
			//fmt.Printf("******sub = %v i = %v lenght = %v start = %v tmpCounter = %v totalLen = %v\n", s[i:i+length], i, length, start, tmpCounter, totalLen)
			if checkWords(tmpCounter) && (i+length-start == totalLen) {
				res = append(res, start)
				continue
			}
			i = i + length - 1
		} else {
			start++
			i = start - 1
			tmpCounter = copyMap(counter)
		}
	}
	return res
}

func checkWords(s map[string]int) bool {
	flag := true
	for _, v := range s {
		if v > 0 {
			flag = false
			break
		}
	}
	return flag
}

func copyMap(s map[string]int) map[string]int {
	c := map[string]int{}
	for k, v := range s {
		c[k] = v
	}
	return c
}
