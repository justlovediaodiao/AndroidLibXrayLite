package libv2ray

import (
	"testing"
	"time"
)

type supportSet struct{}

func (s *supportSet) Setup(Conf string) int {
	return 0
}

func (s *supportSet) Prepare() int {
	return 0
}

func (s *supportSet) Shutdown() int {
	return 0
}

func (s *supportSet) Protect(int) bool {
	return true
}

func (s *supportSet) OnEmitStatus(int, string) int {
	return 0
}

func TestMain(t *testing.T) {
	p := V2RayPoint{
		SupportSet:           &supportSet{},
		ConfigureFileContent: `{"listen": "127.0.0.1:10808", "protocol": "socks", "server": "96.45.190.54:443", "password": "FuckGFW10000Times", "cert": "-----BEGIN CERTIFICATE-----\nMIICLTCCAdSgAwIBAgIUNgqvdAGPsQ+t4ctIzFgtVwRtXGQwCgYIKoZIzj0EAwIwfTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDzANBgNVBAoMBkdpdEh1YjEZMBcGA1UECwwQanVzdGxvdmVkaWFvZGlhbzEVMBMGA1UEAwwMOTYuNDUuMTkwLjU0MB4XDTIzMDMyMDA5MjEzMVoXDTMzMDMxNzA5MjEzMVowfTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExFjAUBgNVBAcMDVNhbiBGcmFuY2lzY28xDzANBgNVBAoMBkdpdEh1YjEZMBcGA1UECwwQanVzdGxvdmVkaWFvZGlhbzEVMBMGA1UEAwwMOTYuNDUuMTkwLjU0MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEI8ThVHYKTIuvr3UvROOlkuJYoDpzNBGyl8eUDg1TESkGUv0qBa54LKi3AwAHlLpVROyEZuK7O8Jp8EM+gJ2RlaMyMDAwDwYDVR0RBAgwBocEYC2+NjAdBgNVHQ4EFgQUbYp/ATg/vuEXK9Xv0xineCMSnh8wCgYIKoZIzj0EAwIDRwAwRAIgePGnv8u0OKrsPkGK4Hh9c6uOl9VIqy0GRtpr8Q9a7nICICBOjQda2fQ5DwfJvxoIozexH5o/0OOOv2plAbRH2yqS\n-----END CERTIFICATE-----"}`,
	}
	p.RunLoop(false)
	time.AfterFunc(10*time.Second, func() {
		p.StopLoop()
	})
	<-(chan int)(nil)
}
