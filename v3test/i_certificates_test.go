package v3test

import (
	hwxpay "hwxpay/v3"
	"testing"
)

func TestCertificates(t *testing.T) {
	cert, err := hwxpay.Certificates(config)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(cert)
}
