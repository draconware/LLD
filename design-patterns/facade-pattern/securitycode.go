package main

import "fmt"

type SecurityCode struct {
	code string
}

func newSecurityCode(securityCode string) *SecurityCode {
	return &SecurityCode{
		code: securityCode,
	}
}

func (sc *SecurityCode) checkSecurityCode(code string) error {
	if sc.code == code {
		return nil
	}
	return fmt.Errorf("security code is incorrect")
}
