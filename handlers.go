package main

import (
	"github.com/vladbara705/NeuralNetwork/app/network"
	"net/http"
)

func RunNetworkHandler(w http.ResponseWriter, r *http.Request) {
	test := []float64{1, 0}
	network.Execute(test)
}
