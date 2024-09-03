package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

var endpoint = "http://localhost:8547"

func TestBesuReceipt(t *testing.T) {

	gs := New(endpoint)

	t.Run("Single Receipt", func(t *testing.T) {
		hash := "0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5"

		r, err := gs.GetReceipt(hash)
		if err != nil {
			t.Error(err)
		}

		j, err := json.MarshalIndent(r, "", "\t")
		if err != nil {
			t.Error(err)
		}
		t.Log(string(j))

	})

	t.Run("Multi Receipts", func(t *testing.T) {
		hashList := []string{
			"0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5",
			"0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5",
		}

		r, err := gs.GetMultiReceipt(hashList)
		if err != nil {
			t.Error(err)
		}

		j, err := json.MarshalIndent(r, "", "\t")
		if err != nil {
			t.Error(err)
		}
		t.Log(string(j))
	})

}

func TestBesuReceiptRawForm(t *testing.T) {

	gs := New(endpoint)

	t.Run("Single Receipt", func(t *testing.T) {

		hash := "0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5"
		req := NewRequest(receiptQuery)

		req.Var("txHash", hash)

		// rtn := types.Receipt{

		// }
		var res json.RawMessage

		err := gs.client.Run(context.Background(), req, &res)
		if err != nil {
			fmt.Printf("client run 에러 %s", err.Error())
		} else {
			fmt.Println(string(res))
		}
	})

	t.Run("Multi Receipt", func(t *testing.T) {
		hashList := []string{"0xa4ba3d49a84981185771faa2aa0f5ef2e5466676f1425fd3fbc2fedee3e345f5"}

		var varBuilder strings.Builder
		var queryBuilder strings.Builder

		for i := 0; i < len(hashList); i++ {

			if i != 0 {
				varBuilder.WriteString(", ")
			}
			varBuilder.WriteString(fmt.Sprintf(receiptVariableForm, i))

			queryBuilder.WriteString(fmt.Sprintf(receiptForm, i, i))

		}

		query := fmt.Sprintf(multiReceiptQuery, varBuilder.String(), queryBuilder.String())
		// fmt.Println(query)

		req := NewRequest(query)

		for i := 0; i < len(hashList); i++ {
			req.Var(fmt.Sprintf("txHash%d", i), hashList[i])
		}

		var res json.RawMessage

		err := gs.client.Run(context.Background(), req, &res)
		if err != nil {
			fmt.Printf("client run 에러 %s", err.Error())
		} else {
			fmt.Println(string(res))
		}

	})
}
