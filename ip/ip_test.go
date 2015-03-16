package ip_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kyani-inc/go-utils/ip"
)

func TestIPv4(t *testing.T) {
	expected := "127.0.0.1"

	req, _ := http.NewRequest("GET", "http://example", nil)
	req.Header.Add("X-Forwarded-For", expected)

	addr := ip.Client(req)

	if addr != expected {
		t.Errorf("Wrong IP Address returned. Expected %s but received %s.\n", expected, addr)
	}
}

func TestIPv4_WithPort(t *testing.T) {
	expected := "127.0.0.1"
	actual := fmt.Sprintf("%s:%d", expected, 59160)

	req, _ := http.NewRequest("GET", "http://example", nil)
	req.Header.Add("X-Forwarded-For", actual)

	addr := ip.Client(req)

	if addr != expected {
		t.Errorf("Wrong IP Address returned. Expected %s but received %s.\n", expected, addr)
	}
}

func TestIPv4_FromCDN(t *testing.T) {
	expected := "127.0.0.1"
	actual := fmt.Sprintf("%s, 192.168.0.1:52190, 10.0.0.5", expected)

	req, _ := http.NewRequest("GET", "http://example", nil)
	req.Header.Add("X-Forwarded-For", actual)

	addr := ip.Client(req)

	if addr != expected {
		t.Errorf("Wrong IP Address returned. Expected %s but received %s.\n", expected, addr)
	}
}

func TestIPv6(t *testing.T) {
	expected := "2001:4860:4860::8888"
	actual := fmt.Sprintf("[%s]", expected)

	req, _ := http.NewRequest("GET", "http://example", nil)
	req.Header.Add("X-Forwarded-For", actual)

	addr := ip.Client(req)

	if addr != expected {
		t.Errorf("Wrong IP Address returned. Expected %s but received %s.\n", expected, addr)
	}
}

func TestIPv6_WithPort(t *testing.T) {
	expected := "2001:4860:4860::8888"
	actual := fmt.Sprintf("[%s]:%d", expected, 59160)

	req, _ := http.NewRequest("GET", "http://example", nil)
	req.Header.Add("X-Forwarded-For", actual)

	addr := ip.Client(req)

	if addr != expected {
		t.Errorf("Wrong IP Address returned. Expected %s but received %s.\n", expected, addr)
	}
}

func TestIPv6_FromCDN(t *testing.T) {
	expected := "::1"
	actual := fmt.Sprintf("[%s], [2001:4860:4860::8888]:51950, 10.0.0.1", expected)

	req, _ := http.NewRequest("GET", "http://example", nil)
	req.Header.Add("X-Forwarded-For", actual)

	addr := ip.Client(req)

	if addr != expected {
		t.Errorf("Wrong IP Address returned. Expected %s but received %s.\n", expected, addr)
	}
}
