package refactor

import "fmt"

type Object interface {
	toHTTPRequest() string
}

type Requester interface {
	toHTTPRequest(obj *Object) string
}

type ObjectRequester struct {
	reqr Requester
}

type InboundMaster1 struct {
	MAWB string
	weight string
}

type InboundMaster2 struct {
	MAWB string
	weight string
}

func (ojr *ObjectRequester) toHTTPRequest(obj Object) (string) {
	return obj.toHTTPRequest()
}

func (o *InboundMaster1) toHTTPRequest() (string) {
	var req string = "IM1"
	return req
}
func (o *InboundMaster2) toHTTPRequest() (string) {
	var req string = "IM2"
	return req
}

func Run() {
	var im1 = InboundMaster1{
		MAWB: "MAWB_1",
		weight: "20",
	}
	
	fmt.Println(???.toHTTPRequest(im1)) // error
}