/*
Author:ydy
Date:
Desc:
*/
package testx

import (
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestAdd_UseGoMock(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockDB(ctrl)
	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}

func TestStr(t *testing.T){
	///firstUniqC("abcda")
// var res =	isPalindrome("0P")
	//var res = myAtoi("words and 987")
	//var res = strStr("hello","ll")
	var res = countAndSay(4)
 fmt.Println(res)

}

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
        start := 0
        j := 0
		for j < len(prev) {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			var sss = prev[start]
			fmt.Println(string(sss))
			cur.WriteByte(sss)
			start = j
		}
		prev = cur.String()
	}
	return prev
}

func strStr(haystack string, needle string) int {
	if needle == ""{
		return 0
	}
	var length = len(needle);
	var total = len(haystack) - length + 1;
	for start := 0;start < total;start++ {
		if (haystack[start:start + length] == (needle)) {
			return start;
		}
	}
	return -1;
}

func myAtoi(s string) int {
	var b = []rune(s)
	var c string
	var sign string
	for i:=0;i<len(b);i++{
		if isNumAndOther1(b[i]){
			if isSign(b[i]){
				sign = string(b[i])
			}else{
				c = c + string(b[i])
			}
		}else{
			if isSign(b[i]) && len(c) > 0{
				var result1,ok = strconv.Atoi(c)
				if ok != nil{
					return 0
				}
				return result1
			}
			if len(c) >0 {
				var result1,ok = strconv.Atoi(c)
				if ok != nil{
					return 0
				}
				return result1
			}
			if len(c) == 0 && isNotNum(b[i]){
				return 0
			}
		}
	}
	if len(sign) >0 {
		c = sign + c
	}
	var result,_ = strconv.Atoi(c)
	return result
}

func isNumAndOther1(s1 rune) bool{
	return (s1 >=0 && s1 <= 9 || s1 >= '0' && s1 <= '9' || s1 == '+' || s1 == '-')
}

func isNotNum(s1 rune) bool{
	return (s1 >='a' && s1 <= 'z' || s1 >='A' && s1 <= 'Z' || s1 == '.')
}

func isSign(s1 rune) bool {
	return  s1 == '+' || s1 == '-'
}

func firstUniqC(s string) int {
	var b = []rune(s)
	var bMap map[rune]int = make(map[rune]int)
	var bMapIndex = make(map[rune]int)
	for i:=0;i<len(b);i++{
		var val,ok = bMap[b[i]]
		if ok{
			bMap[b[i]] = val + 1
		}else{
			bMap[b[i]] = 1
			bMapIndex[b[i]] = i
		}
	}
	for k,v := range bMap{
		if v == 1{
			return bMapIndex[k]
		}
	}
	return -1
}

func isPalindrome(s string) bool {
	var s2 string
	for i:=0;i<len(s);i++{
		if isNumAndOther(rune(s[i])){
			s2 = s2 + string(s[i])
		}
	}
	var b = []rune(strings.ToLower(s2))
	var c = len(b)
	for i:= 0;i<c/2;i++{
		if b[i] != b[c-i-1]{
			return false
		}
	}
	return true
}

func isNumAndOther(s1 rune) bool{
	return (s1 >='a' && s1 <= 'z' || s1 >='A' && s1 <= 'Z' || s1 >=0 && s1 <= 9 || s1 >= '0' && s1 <= '9')
}

type Handler struct {

}
var _ Handler = Handler{}
func (h Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
  //errors.Is()
	//fmt.Errorf()
	//errors.Unwrap()
	//signals := make(chan os.Signal,2)
	//signal.Notify(signals,ShutdownSignals...)
}








































































































































