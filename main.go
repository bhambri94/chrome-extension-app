package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	router.HandleFunc("/v1/chrome/extensions/category={category}&users={users}&ratings={ratings}&reviews_count={reviews_count}", handleChromeExtensionWithCategoryAndFilters)
	router.HandleFunc("/v1/chrome/extensions", handleChromeExtensionWithNoFilters)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func handleChromeExtensionWithNoFilters(w http.ResponseWriter, r *http.Request) {
	sugar.Infof("calling chrome extension manager api!")
	configs.SetConfig()
	finalValues, error := chrome.GetChromeExtension("", "", "", "")
	finalValuesCount := len(finalValues)
	if error == nil && finalValuesCount > 0 {
		sugar.Infof("Completed! chrome extension manager api with filters!")
		err := sheets.ClearSheet(configs.Configurations.SheetNameWithRange + ":M120000")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		}
		err = sheets.BatchWrite(configs.Configurations.SheetNameWithRange, finalValues)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode("Something went wrong. Seems like you have breched the max limit of Google Sheets to fill 500000 cells. Please use filters")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Sheet has been updated: https://docs.google.com/spreadsheets/d/12nVbO18suvGe_vnapNh099pFO0r5pexek2fjYdEFOSI/edit#gid=0")
		}
	} else if error != nil {
		err := fmt.Sprintf("%v", error)
		sugar.Infof("Failed! chrome extension manager api failed with error!" + err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed! chrome extension manager api failed with error")
	} else if finalValuesCount == 0 {
		sugar.Infof("Failed! chrome extension manager api is giving 0 results, hence not updating Sheets!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Failed! chrome extension manager api is giving 0 results, Not updating Google Sheets, Please check your request")
	}
}

func handleChromeExtensionFilters(w http.ResponseWriter, r *http.Request) {
	sugar.Infof("calling chrome extension manager api with filters!")
	configs.SetConfig()
	users := mux.Vars(r)["users"]
	ratings := mux.Vars(r)["ratings"]
	reviews_count := mux.Vars(r)["reviews_count"]
	fmt.Println(users, ratings, reviews_count)
	finalValues, error := chrome.GetChromeExtension("", users, ratings, reviews_count)
	finalValuesCount := len(finalValues)
	if error == nil && finalValuesCount > 0 {
		sugar.Infof("Completed! chrome extension manager api with filters!")
		err := sheets.ClearSheet(configs.Configurations.SheetNameWithRange + ":M120000")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
		}
		err = sheets.BatchWrite(configs.Configurations.SheetNameWithRange, finalValues)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode("Something went wrong. Seems like you have breched the max limit of Google Sheets to fill 500000 cells. Please use filters")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Sheet has been updated: https://docs.google.com/spreadsheets/d/12nVbO18suvGe_vnapNh099pFO0r5pexek2fjYdEFOSI/edit#gid=0")
		}
	} else if error != nil {
		err := fmt.Sprintf("%v", error)
		sugar.Infof("Failed! chrome extension manager api failed with error!" + err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Failed! chrome extension manager api failed with error")
	} else if finalValuesCount == 0 {
		sugar.Infof("Failed! chrome extension manager api is giving 0 results, hence not updating Sheets!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Failed! chrome extension manager api is giving 0 results, Not updating Google Sheets, Please check your request")
	}
}

func handleChromeExtensionWithCategoryAndFilters(w http.ResponseWriter, r *http.Request) {
	sugar.Infof("calling chrome extension manager api with category and filters!")
	configs.SetConfig()
	category := mux.Vars(r)["category"]
	category = strings.ToLower(category[1 : len(category)-1])
	users := mux.Vars(r)["users"]
	ratings := mux.Vars(r)["ratings"]
	reviews_count := mux.Vars(r)["reviews_count"]
	sugar.Infof(category, users, ratings, reviews_count)
	finalValues, error := chrome.GetChromeExtension(category, users, ratings, reviews_count)
	finalValuesCount := len(finalValues)
	if error == nil && finalValuesCount > 0 {
		sugar.Infof("Completed! chrome extension manager api with filters!")
		err := sheets.ClearSheet(configs.Configurations.SheetNameWithRange + ":M120000")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode("Something went wrong.")
		}
		err = sheets.BatchWrite(configs.Configurations.SheetNameWithRange, finalValues)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode("Something went wrong. Seems like you have breched the max limit of Google Sheets to fill 500000 cells. Please use filters")
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode("Sheet has been updated: https://docs.google.com/spreadsheets/d/12nVbO18suvGe_vnapNh099pFO0r5pexek2fjYdEFOSI/edit#gid=0")
		}
	} else if error != nil {
		err := fmt.Sprintf("%v", error)
		sugar.Infof("Failed! chrome extension manager api failed with error!" + err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Failed! chrome extension manager api failed with error")
	} else if finalValuesCount == 0 {
		sugar.Infof("Failed! chrome extension manager api is giving 0 results, hence not updating Sheets!")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode("Failed! chrome extension manager api is giving 0 results, Not updating Google Sheets, Please check your request")
	}

}
