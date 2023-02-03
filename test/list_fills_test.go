package test

import (
	"fmt"
	"os"
	"quantfu.com/coinbase-adv/client"
	"quantfu.com/coinbase-adv/model"
	"testing"
	"time"
)

func Test_ListFills(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	p := &client.ListOrdersParams{
		OrderStatus: []string{string(model.FILLED)},
	}

	rsp, err := cln.ListOrders(p)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	// should be at least 1
	if len(rsp.Orders) < 1 {
		t.FailNow()
	}

	for _, o := range rsp.Orders {

		// list fills.
		pf := &client.ListFillsParams{
			OrderId:   *o.OrderId,
			ProductId: *o.ProductId,
		}

		frsp, err := cln.ListFills(pf)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, f := range frsp.Fills {

			tm, _ := time.Parse(time.RFC3339, f.GetTradeTime())

			fmt.Printf("[%s] [%s] [%s] %f  @ %f \n", tm.String(), f.GetProductId(), f.GetSide(), f.GetSize(), f.GetPrice())

		}
	}

}
