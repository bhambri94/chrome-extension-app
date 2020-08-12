package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bhambri94/chrome-extension-app/chrome"
	"github.com/bhambri94/chrome-extension-app/configs"
	"github.com/bhambri94/chrome-extension-app/sheets"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction()
	sugar     = logger.Sugar()
)

func main() {
	sugar.Infof("starting chrome extension manager app server...")
	defer logger.Sync() // flushes buffer, if any

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/chrome/extensions/category={category}&users={users}&ratings={ratings}&reviews_count={reviews_count}", handleChromeExtension)
	router.HandleFunc("/v1/chrome/extensions", handleChromeExtensionWithNoFilters)
	// router.GET("/v1/chrome/extensions", handleChromeExtension)
	// router.GET("/v1/chrome/extensions/category={category?}/users>{users?}/ratings>{ratings?}/reviews_count>{reviews_count?}", handleChromeExtension)
	// router.GET("/hello/{name}", Hello)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleChromeExtensionWithNoFilters(w http.ResponseWriter, r *http.Request) {
	sugar.Infof("calling chrome extension manager api!")
	configs.SetConfig()
	finalValues := chrome.GetChromeExtension()
	sheets.ClearSheet(configs.Configurations.SheetNameWithRange)
	sheets.BatchWrite(configs.Configurations.SheetNameWithRange, finalValues)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Sheet has been successfully updated")
}

func handleChromeExtension(w http.ResponseWriter, r *http.Request) {
	sugar.Infof("calling chrome extension manager api with filters!")
	category := mux.Vars(r)["category"]
	users := mux.Vars(r)["users"]
	ratings := mux.Vars(r)["ratings"]
	reviews_count := mux.Vars(r)["reviews_count"]
	fmt.Println(category, users, ratings, reviews_count)
	sugar.Infof("calling chrome extension manager api!")
	configs.SetConfig()
	finalValues := chrome.GetChromeExtensionWithFilters(category, users, ratings, reviews_count)
	sheets.ClearSheet(configs.Configurations.SheetNameWithRange + ":M120000")
	sheets.BatchWrite(configs.Configurations.SheetNameWithRange, finalValues)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Sheet has been successfully updated")
}
