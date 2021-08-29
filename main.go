//dig +trace

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lixiangzhong/dnsutil"
	"github.com/miekg/dns"
)

func getRecordFromNS(c *gin.Context) {
	recordType := c.DefaultQuery("type", "A")
	nameserver := c.Query("ns")
	domain := c.Query("name")

	if nameserver == "" || domain == "" {
		c.String(http.StatusBadRequest, "")
		return
	}

	rdns, err := dns.ReverseAddr(domain)
	if err == nil {
		domain = rdns
		if recordType == "A" || recordType == "AAAA" {
			recordType = "PTR"
		}
	}

	var dig dnsutil.Dig
	dig.At(nameserver)

	var res interface{}

	switch recordType {
	case "A":
		res, err = dig.A(domain)
	case "AAAA":
		res, err = dig.AAAA(domain)
	case "CAA":
		res, err = dig.CAA(domain)
	case "CNAME":
		res, err = dig.CNAME(domain)
	case "MX":
		res, err = dig.MX(domain)
	case "NS":
		res, err = dig.NS(domain)
	case "PTR":
		res, err = dig.PTR(domain)
	case "SRV":
		res, err = dig.SRV(domain)
	case "TXT":
		res, err = dig.TXT(domain)
	}

	if err != nil {
		c.String(http.StatusBadRequest, "")
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func getTrace(c *gin.Context) {
	domain := c.Query("name")

	var dig dnsutil.Dig
	res, err := dig.Trace(domain)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}

func main() {
    gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/trace", getTrace)
	r.GET("/record", getRecordFromNS)

	r.Run()
}
