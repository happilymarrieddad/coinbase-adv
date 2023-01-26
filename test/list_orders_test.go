package test

import (
	"coinbase-adv/client"
	"coinbase-adv/model"
	"fmt"
	"os"
	"testing"
	"time"
)

func Test_ListOrders(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	rsp, err := cln.ListOrders(nil)
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
		fmt.Printf("[%s] [%s] [%s] %f  @ %f \n", o.GetStatus(), o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
	}

	for rsp.GetHasNext() {

		p := client.ListOrdersParams{
			Cursor: rsp.Cursor,
		}

		rsp, err = cln.ListOrders(&p)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, o := range rsp.Orders {
			fmt.Printf("[%s] [%s] %f  @ %f \n", o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
		}
	}
}

func Test_ListOrdersOpen(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	p := client.ListOrdersParams{
		OrderStatus: []string{string(model.OPEN)},
	}

	rsp, err := cln.ListOrders(&p)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	for _, o := range rsp.Orders {
		fmt.Printf("[%s] [%s] [%s] %f  @ %f \n", o.GetStatus(), o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
	}

	for rsp.GetHasNext() {

		p := client.ListOrdersParams{
			OrderStatus: []string{string(model.OPEN)},
			Cursor:      rsp.Cursor,
		}

		rsp, err = cln.ListOrders(&p)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, o := range rsp.Orders {
			fmt.Printf("[%s] [%s] %f  @ %f \n", o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
		}
	}
}

func Test_ListOrdersFilled(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	p := client.ListOrdersParams{
		OrderStatus: []string{string(model.FILLED)},
	}

	rsp, err := cln.ListOrders(&p)
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
		fmt.Printf("[%s] [%s] [%s] %f  @ %f \n", o.GetStatus(), o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
	}

	for rsp.GetHasNext() {

		p := client.ListOrdersParams{
			OrderStatus: []string{string(model.FILLED)},
			Cursor:      rsp.Cursor,
		}

		rsp, err = cln.ListOrders(&p)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, o := range rsp.Orders {
			fmt.Printf("[%s] [%s] %f  @ %f \n", o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
		}
	}
}

func Test_ListOrdersDateFilter(t *testing.T) {
	//devToken := os.Getenv("CB-ACTOKEN")
	//creds := client.Credentials{AccessToken: devToken}

	creds := client.Credentials{
		ApiKey:      os.Getenv("CB-APIKEY"),
		ApiSKey:     os.Getenv("CB-SKEY"),
		AccessToken: os.Getenv("CB-ACTOKEN"),
	}

	cln := client.NewClient(&creds)

	et := time.Now().UTC()
	st := et.AddDate(0, 0, -14)

	p := client.ListOrdersParams{
		StartDate: st,
		EndDate:   et,
	}

	rsp, err := cln.ListOrders(&p)
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
		fmt.Printf("[%s] [%s] [%s] %f  @ %f \n", o.GetStatus(), o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
	}

	for rsp.GetHasNext() {

		p := client.ListOrdersParams{
			Cursor: rsp.Cursor,
		}

		rsp, err = cln.ListOrders(&p)
		if err != nil {
			fmt.Println(err)
			t.FailNow()
			return
		}

		for _, o := range rsp.Orders {
			fmt.Printf("[%s] [%s] %f  @ %f \n", o.GetProductId(), o.GetSide(), o.GetFilledSize(), o.GetAverageFilledPrice())
		}
	}
}