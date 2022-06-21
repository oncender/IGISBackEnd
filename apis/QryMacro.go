package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ServeMacroWhole(d IDPMacro, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	qry := procMacroParam(params)

	packet, _ := json.Marshal(serveMacroQry(qry, d))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "application/json")
	w.Write(packet)
}

func procMacroParam(m map[string]string) ReqIDPMacro {
	q := ReqIDPMacro{
		YearFrom: func() int {
			yf, _ := strconv.Atoi(m["yearFrom"])
			return yf
		}(),
		YearUntil: func() int {
			yu, _ := strconv.Atoi(m["yearUntil"])
			return yu
		}(),
		Commodity: m["commodity"],
	}
	return q
}

func serveMacroQry(rq ReqIDPMacro, d IDPMacro) []macroRow {
	dMap := map[string][]macroRow{
		"kr1y":      d.Data.KR1Y,
		"kr3y":      d.Data.KR3Y,
		"kr5y":      d.Data.KR5Y,
		"ifd1y":     d.Data.IFD1Y,
		"cd91d":     d.Data.CD91D,
		"cp91d":     d.Data.CP91D,
		"koribor3m": d.Data.KORIBOR3M,
	}
	var sendPacket = []macroRow{}
	for _, row := range dMap[rq.Commodity] {
		y, _ := strconv.Atoi(row.Date[:4])
		if (y >= rq.YearFrom) && (y <= rq.YearUntil) {
			sendPacket = append(sendPacket, row)
		}
	}
	return sendPacket
}
