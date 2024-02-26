package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error ", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, `Goodbye %s`, d)
}
