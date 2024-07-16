package deposit

import (
	"fmt"
	"testing"
)

func TestQueryParams(t *testing.T) {
	a, _ := QueryParams(QueryParamsRequest{})
	fmt.Println(a)
}
