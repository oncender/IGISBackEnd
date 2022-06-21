package apis

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

const (
	DATA_PANIC_ASSET = "Checklist :: Panic :: "
	DATA_PANIC_DEBT  = "Debt :: Panic :: "
	DATA_PANIC_MACRO = "Macro :: Panic :: "

	DATA_ERR_ASSET = "Checklist :: Error :: "
	DATA_ERR_DEBT  = "Debt :: Error :: "
	DATA_ERR_MACRO = "Macro :: Error :: "
)

const (
	MSG_ASSET = "Asset :: "
	MSG_DEBT  = "Debt :: "
	MSG_MACRO = "Macro :: "
)

const (
	TEST_URL_ROW   = "http://localhost:8080/api/v1/debtRowCount?yearFrom=2000&yearUntil=2021&aumFrom=1&aumUntil=100000000000&debtFrom=1&debtUntil=1000000000000"
	TEST_URL_ASSET = "http://localhost:8080/api/v1/asset?strat=Core"
	TEST_URL_DEBT  = "http://localhost:8080/api/v1/debt?yearFrom=2010&yearUntil=2020"
	TEST_URL_MACRO = "http://localhost:8080/api/v1/macro?commodity=kr1y&yearFrom=2010&yearUntil=2020"
)

func ServeLanding(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "IGIS Debt Platform landing page\n")
}

func IsWithInSlider(s, e int, val string) bool {
	v := strings.Replace(val, ",", "", -1)
	res, _ := strconv.Atoi(v)
	if (res >= s) && (res <= e) {
		return true
	} else {
		return false
	}
}

func IsWithInChoice(isSame, val string) bool {
	workString := strings.Replace(isSame, " ", "", -1)
	workStringSlice := strings.Split(workString, "-")
	result := false
	for _, i := range workStringSlice {
		if i == val {
			result = true
		}
	}
	return result
}
