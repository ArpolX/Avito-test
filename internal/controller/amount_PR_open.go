package controller

import (
	"fmt"
	"net/http"
)

func (c *ControllerImpl) AmountPROpen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	n, err := c.Srv.AmountPROpen(ctx)
	if err != nil {
		CreateError("400", err.Error(), w)
		return
	}

	w.Write([]byte(fmt.Sprintf("Количество pr со статусом Open: %v", n)))
}
