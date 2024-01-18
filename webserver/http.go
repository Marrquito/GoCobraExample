package webserver

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Marrquito/GoCobraExample/app"
)

type Server struct {
	Port string
}

func (s Server) SumHandler(w http.ResponseWriter, r *http.Request) {
	a, err := strconv.ParseFloat(r.URL.Query().Get("a"), 64)
	if err != nil {
		a = 0
	}
	
	b, err := strconv.ParseFloat(r.URL.Query().Get("b"), 64)
	if err != nil {
		b = 0
	}

	calc := app.NewCalc()

	calc.A = a
	calc.B = b

	w.Write([]byte(fmt.Sprintf("Resultado da soma: %.2f\n", calc.Sum())))
}

func (s Server) Server() {
	http.HandleFunc("/", s.SumHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}