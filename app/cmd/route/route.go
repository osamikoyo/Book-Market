package route

import (
	"book-market/api"
	"log/slog"
	"net/http"
	"os"
)


func Route()  {
	http.HandleFunc("/register", api.HandleRegister)
	http.HandleFunc("/getbook", api.HandleBookmarketGetBooks)

	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	loger.Info("Server is running!\n")

	
	loger.Error(http.ListenAndServe(":2020", nil).Error())
}