package main

import (
	"net"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func RemoteAddress(c echo.Context) string {
	AddressAndPort := strings.Split(c.Request().RemoteAddr, ":")
	Address := AddressAndPort[0]
	return Address
}

func ShowIP(c echo.Context) error {
	Address := RemoteAddress(c)
	return c.String(http.StatusOK, Address)
}

func ReverseDNS(c echo.Context) error {
	Address := RemoteAddress(c)
	Hostnames, err := net.LookupAddr(Address)
	if err != nil {
		return c.String(http.StatusNotFound, "Unable perform reverse DNS search for address " + Address)
	}
	HostnameString := strings.Join(Hostnames, "\n")
	return c.String(http.StatusOK, HostnameString)


}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/rdns*", ReverseDNS)
	e.GET("/*", ShowIP)
	e.Logger.Fatal(e.Start(":8081"))
}
