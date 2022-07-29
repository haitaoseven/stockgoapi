package common

import (
	"fmt"
	"os"
	"strconv"
)

func Float64(v interface{}) float64 {
	if s, ok := v.(string); ok {
		f64, _ := strconv.ParseFloat(s, 10)
		return f64
	}
	if i, ok := v.(int); ok {
		return float64(i)
	}
	if i32, ok := v.(int32); ok {
		return float64(i32)
	}
	if i64, ok := v.(int64); ok {
		return float64(i64)
	}
	return 0
}

func String(v interface{}) string {
	if i, ok := v.(int); ok {
		return strconv.Itoa(i)
	}
	if i32, ok := v.(int64); ok {
		return strconv.FormatInt(int64(i32), 10)
	}
	if i64, ok := v.(int64); ok {
		return strconv.FormatInt(i64, 10)
	}
	if s, ok := v.(string); ok {
		return s
	}
	if f, ok := v.(float32); ok {
		return strconv.FormatFloat(float64(f), 'f', 10, 32)
	}
	if f64, ok := v.(float64); ok {
		return strconv.FormatFloat(f64, 'f', 10, 64)
	}
	return ""
}

/*
//string到int
int,err:=strconv.Atoi(string)
//string到int64
int64, err := strconv.ParseInt(string, 10, 64)
//string到float32(float64)
float,err := strconv.ParseFloat(string,32/64)

//int到string
string:=strconv.Itoa(int)
//int64到string
string:=strconv.FormatInt(int64,10)
//float到string
string := strconv.FormatFloat(float32, 'E', -1, 32)
string := strconv.FormatFloat(float64, 'E', -1, 64)
// 'b' (-ddddp±ddd，二进制指数)
// 'e' (-d.dddde±dd，十进制指数)
// 'E' (-d.ddddE±dd，十进制指数)
// 'f' (-ddd.dddd，没有指数)
// 'g' ('e':大指数，'f':其它情况)
// 'G' ('E':大指数，'f':其它情况)
*/

func Int(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return result
}

func Round(v interface{}, num int) float64 {
	n := strconv.Itoa(num)
	f := "%." + n + "f"
	ret := ""
	if s, ok := v.(string); ok {
		f64, _ := strconv.ParseFloat(s, 10)
		ret = fmt.Sprintf(f, f64)
	}
	if f64, ok := v.(float64); ok {
		ret = fmt.Sprintf(f, f64)
	}
	return Float64(ret)
}

func RoundStr(v interface{}, num int) string {
	ret := Round(v, num)
	return String(ret)
}

func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

func ArrayGtNum(arr []float64, num float64) (ret []float64) {
	ret = []float64{}
	for _, v := range arr {
		if v > num {
			ret = append(ret, v)
		}
	}
	return
}

func ArrayGtEqNum(arr []float64, num float64) (ret []float64) {
	ret = []float64{}
	for _, v := range arr {
		if v >= num {
			ret = append(ret, v)
		}
	}
	return
}

func ArrayLtNum(arr []float64, num float64) (ret []float64) {
	ret = []float64{}
	for _, v := range arr {
		if v < num {
			ret = append(ret, v)
		}
	}
	return
}

func ArrayLteqNum(arr []float64, num float64) (ret []float64) {
	ret = []float64{}
	for _, v := range arr {
		if v <= num {
			ret = append(ret, v)
		}
	}
	return
}

func ArraySum(arr []float64) (sum float64) {
	for _, v := range arr {
		sum = sum + v
	}
	return
}

func ArrayAvg(arr []float64) (avg float64) {
	n := Float64(len(arr))
	var sum float64 = 0
	for _, v := range arr {
		sum = sum + v
	}
	if n == 0 {
		avg = 0
	} else {
		avg = sum / n
	}
	return avg
}

func Max(arr []float64) (max float64) {
	for k, v := range arr {
		if k == 0 {
			max = v
			continue
		}
		if v > max {
			max = v
		}
	}
	return
}

func Min(arr []float64) (min float64) {
	for k, v := range arr {
		if k == 0 {
			min = v
			continue
		}
		if v < min {
			min = v
		}
	}
	return
}

func ArrayReverse(arr []float64) []float64 {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func ArraySliceLast(arr []float64, last int) []float64 {
	if len(arr)-1 > last {
		num := len(arr) - 1 - last
		return arr[num:]
	} else {
		return arr
	}

}
