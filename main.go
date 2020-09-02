package main

//import "strings"

func main() {
	ans := make([]string, 0)
	ans = process([]rune("?01?"), 0, ans)
}

func process(r []rune, i int, ans []string) []string {

	if i == len(r) {
		ans = append(ans, string(r))
	}

	if i > len(r)-1 {
		return ans
	}

	if r[i] == '?' {
		for char := '0'; char <= '1'; char++ {
			r[i] = char
			ans = process(r, i+1, ans)
			r[i] = '?'
		}
		return ans
	}
	return process(r, i+1, ans)

}

/*
0? => [01, 00]
00? => [001?, 000]
0?0? => [0101, 0001, 0100, 0000]
*/
