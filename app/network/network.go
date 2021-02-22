package network

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/vladbara705/NeuralNetwork/heplers"
	"log"
)

var cfg *ini.File

// NeuralNetwork type
type NeuralNetwork struct {
	InputNeurons  int
	HiddenLayers  int
	HiddenNeurons int
	OutputNeurons int
	UseBias 	  bool
	NetworkSettingPath string
}
var NeuralNetworkSetting = &NeuralNetwork{}

func init() {
	var err error
	cfg, err = ini.Load("app/network/network.ini")
	if err != nil {
		log.Fatalf("Ошибка при парсинге файла конфигурации: %v", err)
	}
	heplers.MapTo(cfg, "neuralNetwork", NeuralNetworkSetting)
}



func Execute(data []float64) {
	if len(data) != NeuralNetworkSetting.InputNeurons {
		log.Fatalf("Количество входных параметров не соответствует заявленным")
	}

	/**************
	**	INPUT LAYER
	***************/

	fmt.Println(NeuralNetworkSetting)
}