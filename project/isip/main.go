package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func main() {
	isCidr()
	validIPs := []string{
		"192.168.15.234/24",
		"10.0.0.1/8",
		"172.16.0.0/12",
	}

	invalidIPs := []string{
		"192.168.15.256/24",
		"10.0.0/8",
		"172.16.0.0.0/12",
	}

	for _, ip := range validIPs {
		if isValidIPWithPrefix(ip) {
			fmt.Printf("%s is a valid IP address with prefix\n", ip)
		} else {
			fmt.Printf("%s is an invalid IP address with prefix\n", ip)
		}
	}

	for _, ip := range invalidIPs {
		if isValidIPWithPrefix(ip) {
			fmt.Printf("%s is a valid IP address with prefix\n", ip)
		} else {
			fmt.Printf("%s is an invalid IP address with prefix\n", ip)
		}
	}
}

func isValidIPWithPrefix(ipStr string) bool {
	ipv4Regex := `^(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(3[0-2]|[12]?[0-9])$`
	ipRegex := ipv4Regex
	return regexp.MustCompile(ipRegex).MatchString(ipStr)
}

type MyStruct struct {
	CIDR string `validate:"cidr"`
}

func isCidr() {
	validate := validator.New()

	// 测试合法的 CIDR 格式
	validCIDR := MyStruct{CIDR: "192.168.1.0/24"}
	err := validate.Struct(validCIDR)
	if err != nil {
	} else {
		fmt.Println("Valid CIDR:", validCIDR.CIDR)
	}

	// 测试非法的 CIDR 格式
	invalidCIDR := MyStruct{CIDR: "192.168.1.256/24"}
	err = validate.Struct(invalidCIDR)
	if err != nil {
		fmt.Println("Invalid CIDR:", invalidCIDR.CIDR)
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Field(), err.Tag(), err.Kind())
		}
	} else {
		fmt.Println("Valid CIDR:", invalidCIDR.CIDR)
	}
}
