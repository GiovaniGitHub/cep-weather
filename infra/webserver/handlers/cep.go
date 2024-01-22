package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/GiovaniGitHub/cpf-weather/internal/entity"
)

// func GetAddress(w http.ResponseWriter, r *http.Request) {
// 	cep := r.URL.Path[len("/cep/"):]
// 	ch := make(chan entity.AnyAddress, 2)

// 	if !regexp.MustCompile(`^\d{5}-\d{3}$`).MatchString(cep) {
// 		cep = regexp.MustCompile(`[^\d]`).ReplaceAllString(cep, "")
// 		if len(cep) == 8 {
// 			cep = cep[:5] + "-" + cep[5:]
// 		}

// 	}
// 	go func() {
// 		resp, err := http.Get("https://cdn.apicep.com/file/apicep/" + cep + ".json")
// 		fmt.Printf(resp.Request.URL.String())
// 		if err != nil {
// 			fmt.Println("Erro ao fazer request para API 1:", err)
// 			return
// 		}
// 		defer resp.Body.Close()
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println("Erro ao ler resposta da API 1:", err)
// 			return
// 		}

// 		var addressApiCep entity.AddressApiCep
// 		err = json.Unmarshal(body, &addressApiCep)
// 		if err != nil {
// 			fmt.Println("Erro ao fazer parse da resposta da API 1:", err)
// 			return
// 		}
// 		addressApiCep.Api = "https://cdn.apicep.com/file/apicep/"
// 		ch <- entity.AnyAddress(addressApiCep)
// 	}()

// 	go func() {
// 		resp, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
// 		if err != nil {
// 			fmt.Println("Erro ao fazer request para API 2:", err)
// 			return
// 		}
// 		defer resp.Body.Close()
// 		body, err := io.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println("Erro ao ler resposta da API 2:", err)
// 			return
// 		}

// 		var addressViaCep entity.AddressViaCep
// 		err = json.Unmarshal(body, &addressViaCep)
// 		if err != nil {
// 			fmt.Println("Erro ao fazer parse da resposta da API 2:", err)
// 			return
// 		}
// 		addressViaCep.Api = "http://viacep.com.br/ws/"
// 		ch <- entity.AnyAddress(addressViaCep)
// 	}()

// 	select {
// 	case res := <-ch:
// 		data, err := json.MarshalIndent(res, "", "  ")
// 		if err != nil {
// 			log.Fatalf("Erro ao serializar os dados em JSON: %v", err)
// 		}
// 		log.Print(string(data))
// 		json.NewEncoder(w).Encode(res)

// 	case <-time.After(1 * time.Second):
// 		fmt.Println("Timeout atingido. Nenhuma resposta recebida a tempo.")
// 		http.Error(w, "Timeout atingido. Nenhuma resposta recebida a tempo.", http.StatusRequestTimeout)
// 	}
// }

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
	// if !regexp.MustCompile(`^\d{5}-\d{3}$`).MatchString(cep) {
	// 	cep = regexp.MustCompile(`[^\d]`).ReplaceAllString(cep, "")
	// 	if len(cep) == 8 {
	// 		cep = cep[:5] + "-" + cep[5:]
	// 	}
	// }

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
		fmt.Println("Erro ao ler resposta da API 2:", err)
		return
	}

	var address entity.Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		fmt.Println("Erro ao fazer parse da resposta da API 2:", err)
		return
	}
	if address.Erro {
		http.Error(w, "can not found zipcode", http.StatusNotFound)
		return
	} else {
		result := strings.ReplaceAll(address.Localidade, " ", "+")

		resp, err := http.Get("https://wttr.in/" + result + "?format=j1")
		if err != nil {
			http.Error(w, "can not found zipcode", http.StatusUnprocessableEntity)
			return
		}

		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		var weatherData entity.WeatherData
		err = json.Unmarshal(body, &weatherData)
		if err != nil {
			fmt.Println("Erro ao fazer parse da resposta da API 2:", err)
			return
		}

		var temperature entity.Temperature
		temperature.TempC = weatherData.CurrentCondition[0].TempC
		temperature.TempF = weatherData.CurrentCondition[0].TempF
		value, err := strconv.ParseFloat(weatherData.CurrentCondition[0].TempC, 32)
		if err != nil {
			fmt.Println("Error ao converter o valor para float:", err)
			return
		}
		temperature.TempK = strconv.FormatFloat(value+273.15, 'f', 2, 32)
		json.NewEncoder(w).Encode(temperature)
	}

}
