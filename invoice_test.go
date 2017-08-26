package harvest

import (
	"testing"
)

func TestCSVLineItems(t *testing.T) {
	i := Invoice{
		CSVLineItems: "kind,description,quantity,unit_price,amount,taxed,taxed2,project_id\nProduct,A description,1.00,100.00,100.0,false,false,\n",
	}
	li := i.LineItems()
	if len(li) != 1 {
		t.Fatal("Unexpected length", len(li))
	}
	if li[0]["description"] != "A description" {
		t.Error("Unexpected description", li[0]["description"])
	}
}
