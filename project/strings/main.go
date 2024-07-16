package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	secure, err := RandStringBytesSecure(10)
	if err != nil {
		return
	}
	fmt.Println(secure)
	strSlice()
	var str = `{
    "code": 200,
    "requestId": "121ad273-f246-4b4c-8743-dbc0f333e716",
    "data": [
        {
            "name": "",
            "type": 1,
            "totalRows": 0,
            "children": [
                {
                    "name": "Device",
                    "type": 2,
                    "totalRows": 0,
                    "children": [
                        {
                            "name": "DeviceList",
                            "method": "GET",
                            "api": "ecos-hub/v1/devices",
                            "type": 3,
                            "totalRows": 5,
                            "successRows": 1,
                            "successRate": "20.00%"
                        },
                        {
                            "name": "DeviceList",
                            "method": "POST",
                            "api": "ecos-hub/v1/devices",
                            "type": 3,
                            "totalRows": 1,
                            "successRows": 1,
                            "successRate": "100.00%"
                        }
                    ]
                }
            ]
        }
    ]
}`
	fmt.Println(str)
	str = strings.Replace(str, "\n", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\t", "", -1)

	str = strings.Replace(str, " ", "", -1)
	fmt.Println(str)

}

func strSlice() {
	var str = "{'antiReflux': False, 'backFlowLimitPower': 0, 'chargeInAppointTime': False, 'controlPower': 0, 'demandPower': 0, 'dischargeInAppointTime': False, 'pccDemandPower': 0, 'powerFactorControl': False, 'powerFactorControlValue': 0.9, 'powerFactorFirst': False, 'soc': 10, 'offGirdSoc': 10, 'gridControlPower': 0, 'monthControlPower': 0} "
	str = strings.Replace(str, "'", "", -1)
}

func RandStringBytesSecure(n int) (string, error) {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b)[:n], nil
}
