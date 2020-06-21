package main

import "fmt"

//题目描述:
//将一个字符串中的空格替换成 "%20"
//Input:
//"A B"
//
//Output:
//"A%20B"

//解题思路：

// ① 在字符串尾部填充任意字符，使得字符串的长度等于替换之后的长度。因为一个空格要替换成三个字符（%20），所以当遍历到一个空格时，需要在尾部填充两个任意字符。
//
// ② 令 P1 指向字符串原来的末尾位置，P2 指向字符串现在的末尾位置。P1 和 P2 从后向前遍历，当 P1 遍历到一个空格时，就需要令 P2 指向的位置依次填充 02%（注意是逆序的），否则就填充上 P1 指向字符的值。从后向前遍是为了在改变 P2 所指向的内容时，不会影响到 P1 遍历原来字符串的内容。
//
// ③ 当 P2 遇到 P1 时（P2 <= P1），或者遍历结束（P1 < 0），退出。

func main() {
	s := "ab d  32 x"
	fmt.Printf("Origin string = %v\n", s)
	result := replaceSpace(s)
	fmt.Printf("Replace string = %v", result)
}

func replaceSpace(s string) string {
	//按照空格的数目将原来的字符串长度增加 len(空格数)*2，新增加的位置赋x
	r := []rune(s)
	p1 := len(r) - 1
	for i := 0; i <= p1; i++ {
		if r[i] == ' ' {
			r = append(r, 'x', 'x')
		}
	}

	//p2从尾部追赶p1，
	// 当p1位置为空格时，p2连续前进3个位置并依次赋值 0 2 %
	// 当p1位置为其他字符时，将其直接赋给p2当前的位置。
	p2 := len(r) - 1
	for p1 >= 0 && p2 > p1 {
		if r[p1] == ' ' {
			r[p2] = '0'
			p2--
			r[p2] = '2'
			p2--
			r[p2] = '%'
			p2--
		} else {
			r[p2] = r[p1]
			p2--
		}
		p1--
	}
	return string(r)
}
