package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/GiovaniGitHub/cep-weather/internal/entity"
)

// Get Address godoc
// @Summary      Get Address 2
// @Description  Get Address by Post Code 2
// @Tags         address
// @Accept       json
// @Produce      json
// @Param        cep     path       string    true    "Cep"
// @Success      200     {object}   entity.Temperature
// @Failure      400     string     "Bad Request"
// @Failure      404     string     "Not Found"
// @Failure      500     string     "Internal Server Error"
// @Router       /cep/{cep}	[get]
func GetTemperature(w http.ResponseWriter, r *http.Request) {

	cep := r.URL.Path[len("/cep/"):]

	resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")

	if err != nil {
		http.Error(w, "Erro ao acessar o viacep:", resp.StatusCode)
		return
	}
	if resp.StatusCode == 400 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta da API http://viacep.com.br/ws/:", err)
		return
	}

	var address entity.Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Println("Erro ao fazer parse da resposta da API http://viacep.com.br/ws/:", err)
		return
	}
	if address.Erro {
		http.Error(w, "can not found zipcode", http.StatusNotFound)
		return
	} else {
		result := strings.ReplaceAll(address.Localidade, " ", "+")

		resp, err := http.Get("http://wttr.in/" + result + "?format=j1")
		if err != nil {
			http.Error(w, "can not found zipcode", http.StatusUnprocessableEntity)
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		var weatherData entity.WeatherData
		err = json.Unmarshal(body, &weatherData)
		if err != nil {
			fmt.Println("Erro ao fazer parse da resposta da API https://wttr.in/:", err)
			return
		}

		var temperature entity.Temperature
		temperature.TempC = weatherData.CurrentCondition[0].TempC
		value, err := strconv.ParseFloat(weatherData.CurrentCondition[0].TempC, 32)
		if err != nil {
			fmt.Println("Error ao converter o valor para float:", err)
			return
		}
		temperature.TempF = strconv.FormatFloat(value*1.8+32, 'f', 2, 32)
		temperature.TempK = strconv.FormatFloat(value+273, 'f', 2, 32)
		json.NewEncoder(w).Encode(temperature)
	}

}
