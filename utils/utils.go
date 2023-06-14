/*
	auth: satanfire
	date: 2023/6/14
	desc: utils
*/

package utils

// 下划线转驼峰
func UnderLineToCamel(str string) string {
	res := make([]byte, 0, len(str))
	for i := 0; i < len(str); i++ {
		if i == 0 && str[i] >= 'a' && str[i] <= 'z' {
			res = append(res, str[i]-32)
		} else if str[i] == '_' {
			if i+1 < len(str) && str[i+1] >= 'a' && str[i+1] <= 'z' {
				res = append(res, str[i+1]-32)
				i++
			}
		} else {
			res = append(res, str[i])
		}
	}
	return string(res)
}
