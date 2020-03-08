package handlers

import (
	//"encoding/json"
	"fmt"
	"net/http"

	"github.com/jschaare/kanji-browser/api/internal/utils"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("example request")
	utils.SendResponse(w, http.StatusOK, map[string]string{"message": "hello world"})
}
