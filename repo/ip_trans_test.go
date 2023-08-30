/*
	auth: satanfire
	date: 2023/8/30
	desc: ip address translate
*/

package repo

import (
	"testing"
)

func TestToInt(t *testing.T) {
	type ipPairType struct {
		ip  string
		num int
	}
	cases := []ipPairType{
		ipPairType{
			ip:  "123.255.111.20",
			num: 2080337684,
		},
		ipPairType{
			ip:  "10.0.1.6",
			num: 167772422,
		},
		ipPairType{
			ip:  "127.0.0.1",
			num: 2130706433,
		},
	}
	for _, item := range cases {
		res, err := IpTransIns.ToInt(item.ip)
		if err != nil {
			t.Errorf("ip error, %s\n", err.Error())
			return
		}

		t.Logf("ip:%s, num:%d, res:%d\n", item.ip, item.num, res)
		if res != item.num {
			t.Errorf("trnas error.\n")
			return
		}
		t.Logf("success.\n")
	}
}
