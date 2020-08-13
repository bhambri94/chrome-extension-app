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
	router.HandleFunc("/v1/chrome/extensions/users={users}&ratings={ratings}&reviews_count={reviews_count}", handleChromeExtensionFilters)
	router.HandleFunc("/v1/chrome/extensions/category={category}&users={users}&ratings={ratings}&reviews_count={reviews_count}", handleChromeExtension)
	router.HandleFunc("/v1/chrome/extensions", handleChromeExtensionWithNoFilters)
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
	json.NewEncoder(w).Encode("Sheet has been successfully updated: https://docs.google.com/spreadsheets/d/12nVbO18suvGe_vnapNh099pFO0r5pexek2fjYdEFOSI/edit#gid=0")
}

func handleChromeExtensionFilters(w http.ResponseWriter, r *http.Request) {
	sugar.Infof("calling chrome extension manager api with filters!")
	configs.SetConfig()
	users := mux.Vars(r)["users"]
	ratings := mux.Vars(r)["ratings"]
	reviews_count := mux.Vars(r)["reviews_count"]
	fmt.Println(users, ratings, reviews_count)
	finalValues := chrome.GetChromeExtensionWithoutCategory(users, ratings, reviews_count)
	sugar.Infof("Completed! chrome extension manager api with filters!")
	sheets.ClearSheet(configs.Configurations.SheetNameWithRange + ":M120000")
	sheets.BatchWrite(configs.Configurations.SheetNameWithRange, finalValues)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Sheet has been updated: https://docs.google.com/spreadsheets/d/12nVbO18suvGe_vnapNh099pFO0r5pexek2fjYdEFOSI/edit#gid=0")
}

func handleChromeExtension(w http.ResponseWriter, r *http.Request) {
	sugar.Infof("calling chrome extension manager api with filters!")
	configs.SetConfig()
	category := mux.Vars(r)["category"]
	category = category[1 : len(category)-1]
	users := mux.Vars(r)["users"]
	ratings := mux.Vars(r)["ratings"]
	reviews_count := mux.Vars(r)["reviews_count"]
	fmt.Println(category, users, ratings, reviews_count)
	finalValues := chrome.GetChromeExtensionWithFilters(category, users, ratings, reviews_count)
	sugar.Infof("Completed! chrome extension manager api with filters!")
	sheets.ClearSheet(configs.Configurations.SheetNameWithRange + ":M120000")
	sheets.BatchWrite(configs.Configurations.SheetNameWithRange, finalValues)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Sheet has been updated: https://docs.google.com/spreadsheets/d/12nVbO18suvGe_vnapNh099pFO0r5pexek2fjYdEFOSI/edit#gid=0")
}
