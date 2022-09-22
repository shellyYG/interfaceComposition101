package main
import (
	// "fmt"
	"interface/composition101/refactor"
	// "interface/composition101/example"
)

type Object interface {
	ToHTTPRequest() (string) // string to mimic *http.Request
}

type InboundMaster1 struct {
	MAWB string
	weight string
}

var im1 = InboundMaster1{
	MAWB: "MAWB_1",
	weight: "20",
}

func (o *InboundMaster1) ToHTTPRequest() (string) {
	var req string = "IM1"
	return req
}

type InboundMaster2 struct {
	MAWB string
	weight string
}

var im2 = InboundMaster2{
	MAWB: "MAWB_2",
	weight: "20",
}

func (o *InboundMaster2) ToHTTPRequest() (string) {
	var req string = "IM2"
	return req
}

func main() {
	// fmt.Println(im1.ToHTTPRequest())
	// fmt.Println(im2.ToHTTPRequest())
	refactor.Run()
	// example.Run()
}