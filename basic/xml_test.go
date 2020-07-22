package basic

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
)

func TestExampleXmlParse(t *testing.T) {
	var xmlModel = new(PaymentService)
	var xmlStr = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE paymentService PUBLIC "-//WorldPay//DTD WorldPay PaymentService v1//EN" "http://dtd.worldpay.com/paymentService_v1.dtd">
<paymentService version="1.4" merchantCode="BYTEMODOLUSD">
    <notify>
        <orderStatusEvent orderCode="1558536470120911vioy">
        </orderStatusEvent>
    </notify>
</paymentService>`
	bReader := bytes.NewBufferString(xmlStr)
	decoder := xml.NewDecoder(bReader)
	err := decoder.Decode(xmlModel)
	fmt.Println(err, xmlModel)
}

type PaymentService struct {
	PaymentService string `xml:"paymentService" json:"paymentService"`
	Version        string `xml:"version,attr" json:"version"`
	MerchantCode   string `xml:"merchantCode,attr" json:"merchantCode"`
	Notify         Notify `xml:"notify" json:"notify"`
}
type Notify struct {
	OrderStatusEvent OrderStatusEvent `xml:"orderStatusEvent" json:"orderStatusEvent"`
}
type OrderStatusEvent struct {
	OrderCode string `xml:"orderCode,attr" json:"orderCode"`
}
