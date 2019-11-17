package main

import (
	"fmt"
	"math/rand"
)

const MAXCARD = 0x37

var mj = []int8{
	0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09,

	0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19,

	0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, // 1-9条

	0x31, 0x32, 0x33, 0x34, 0x35, 0x36, MAXCARD, //东西南北中发白
}

type chai struct {
	a, b, c int8
	eye     bool
}

var hsItem = []chai{}

var hsItem1 = []chai{
	{1, 1, 1, false},
}

var hsItem2 = []chai{
	{2, 0, 0, true},
	{2, 2, 2, false},
}

var hsItem3 = []chai{
	{3, 0, 0, false},
	{3, 1, 1, true},
}

var hsItem4 = []chai{
	{4, 1, 1, false},
	{4, 2, 2, true},
}

func getHandCard(cards []int8, n int) []int8 {
	randLen := len(cards)
	fmt.Println("cards len = ", randLen)
	var ret = make([]int8, MAXCARD+1)
	for index := 0; index < n; index++ {
		rn := rand.Intn(randLen)
		ret[cards[rn]]++
		cards[rn], cards[randLen-1] = cards[randLen-1], cards[rn]
		randLen--
	}
	return ret
}

func checkZi(p []int8, eye *bool) bool {
	for index := 0; index < len(p); index++ {
		if p[index] == 1 || p[index] == 4 {
			return false
		}
		if p[index] == 2 && *eye == true {
			return false
		}
		if p[index] == 2 {
			*eye = true
		}
	}
	return true
}

func getCardCount(p []int8) int8 {
	var sum int8 = 0
	for index := 0; index < len(p); index++ {
		sum += p[index]
	}
	return sum
}

func getChaiMethod(n int8) []chai {
	switch n {
	case 1:
		return hsItem1
	case 2:
		return hsItem2
	case 3:
		return hsItem3
	case 4:
		return hsItem4
	default:
		return hsItem
	}
}

func chaiPai(p []int8, eye *bool) bool {
	for index := 0; index < len(p); index++ {
		if p[index] == 0 {
			continue
		}
		hs := getChaiMethod(p[index])
		for i := 0; i < len(hs); i++ {
			if index+1
			
			if hs[i].b > p[index+1] || hs[i].c > p[index+2] {
				continue
			}
			if hs[i].eye && *eye {
				continue
			}
			p[index] = 0
			p[index+1] -= hs[i].b
			p[index+2] -= hs[i].c
			chaiPai(p[1:], eye)
			p[index] = hs[i].a
			p[index+1] += hs[i].b
			p[index+2] += hs[i].c

		}
	}
}

func checkColor(p []int8, eye *bool) bool {
	cardCnt := getCardCount(p)
	if cardCnt == 0 {
		return true
	}
	if cardCnt%3 == 1 {
		return false
	}
	if cardCnt%3 == 2 && *eye {
		return false
	}

	return chaiPai(p, eye)
}

func isHu(cards []int8) bool {
	var eye = false
	if !checkZi(cards[0x31:], &eye) {
		return false
	}
	if !checkColor(cards[0x01:0x10], &eye) {
		return false
	}
}

func main() {
	var cards []int8
	for index := 0; index < 4; index++ {
		cards = append(cards, mj...)
	}
	handCard := getHandCard(cards, 14)
	fmt.Println("handcard : ", handCard, " , size = ", len(handCard))
	isHu(handCard)
	fmt.Println("hello world")
}
