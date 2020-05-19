package main

import (
	"net/http"

	aero "github.com/aerospike/aerospike-client-go"
)

const (
	// csv-realted
	namespace = "test"
	setName   = "creditcard"

	// BinMap field names
	setNameBin = "set_name"
	ccIDBin    = "ID"
	amountCol  = "AmountBin"
	classBin   = "ClassBin"
	timeBin    = "TimeBin"

	localhost         = "127.0.0.1"
	mlModelServingURL = "http://" + localhost + ":8501/v1/models/fraud:predict"
	inputLength       = 29
	fraudThreshold    = 0.5
)

var aeroClient *aero.Client

// webTxn is a struct for txn incoming in a web request
type webTxn struct {
}

// enrichedTxn is a struct for sending to the ML model
type enrichedTxn struct {
}

// prediction is a struct for gettingt the prediction from the ML model
type prediction struct {
}

// predictionHandler is the entry point to the system,
// ends in validating the prediction
func predictionHandler(w http.ResponseWriter, req *http.Request) {
	// read t§xn, decode JSON, store in Aerospike
	incomingTxn := webTxn{}
	err := acceptTxn(req, aeroClient, &incomingTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// read txn by userID
	enrichedTxn := enrichedTxn{}
	txnOutcome, err := enrichTxn(aeroClient, &incomingTxn, &enrichedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// send enriched txn to model serving web service
	modelPrediction, err := getPrediction(&enrichedTxn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// compare prediction with classification
	validatePrediction(txnOutcome, modelPrediction)
}

// acceptTxn reads an incoming txn and stores it in Aerospike
func acceptTxn(req *http.Request, client *aero.Client, incomingTxn *webTxn) (err error) {
	// read and decode txn

	// store the incoming txn  in Aerospike

	return nil
}

// enrichTxn creates the enriched txn based on the given UserID
func enrichTxn(aeroClient *aero.Client, incomingTxn *webTxn, enrichedTxn *enrichedTxn) (txnOutcome string, err error) {
	// read enriched data by UserID

	// marshal read bins to the enriched struct
	// first build the inner array: [log(amount),v1,...v28]

	// next populate the 2D array

	return "", err
}

// getPrediction sends the enriched txn to the model and gets prediction
func getPrediction(enrichedTxn *enrichedTxn) (modelPrediction string, err error) {
	// prepare the request to the web service serving the model

	// make the request

	// read the response

	// unmarshal to the response struct

	// check the threshold to decide whether it's a fraud

	return "", nil
}

// validatePrediction compares the model prediction with the classification from the DB
func validatePrediction(txnOutcome, modelPrediction string) {
	// compare both predictions

	// advanced: run comparison for all fields in csv
	return
}

func main() {
	// set up a single instance of an Aerospike client
	// connection, it handles the connection pool internally

	// listen and serve
	http.HandleFunc("/", predictionHandler)
	http.ListenAndServe(":8090", nil)
}
