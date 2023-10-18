package cne_express

import (
	"fmt"
	"regexp"
	"shipment-tracker-go/registry"
)

type Cne_express struct {
	IDRegex1 *regexp.Regexp
	/*
		 	IDRegex2 *regexp.Regexp
			IDRegex3 *regexp.Regexp
			IDRegex4 *regexp.Regexp
	*/
}

func (c *Cne_express) matches(input string) bool {
	return c.IDRegex1.MatchString(input) /* || c.IDRegex2.MatchString(input) || c.IDRegex3.MatchString(input) || c.IDRegex4.MatchString(input) */
}

func init() {
	fmt.Println("init")
	registry.Register("cne_express", &Cne_express{
		IDRegex1: regexp.MustCompile(`^3A\d[A-Za-z][0-9]{9}$`),
	})
}
