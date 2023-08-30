/*
	auth: satanfire
	date: 2023/8/30
	desc: ip address translate
*/

package repo

import (
	"errors"
	"strconv"
	"strings"
)

type ipTransIns struct {
}

var IpTransIns *ipTransIns

// ip => int
func (obj *ipTransIns) ToInt(ip string) (int, error) {
	// 拆分ip地址
	ipNums := strings.Split(ip, ".")
	if len(ipNums) != 4 {
		return -1, errors.New("ip is invalid.")
	}

	ipNum, bit := 0, 0
	for i := len(ipNums) - 1; i > -1; i-- {
		tmpNum, err := strconv.Atoi(ipNums[i])
		if err != nil {
			return -1, err
		}

		ipNum += tmpNum * (1 << bit)
		bit += 8
	}
	return ipNum, nil
}

func init() {
	IpTransIns = new(ipTransIns)
}
