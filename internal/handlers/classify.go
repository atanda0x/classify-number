package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/atanda0x/classify-number/internal/funfact"
	"github.com/atanda0x/classify-number/internal/mathutils"
)

type ApiResponse struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	FunFact    string   `json:"fun_fact"`
}

type APIError struct {
	Number string `json:"number"`
	Error  bool   `json:"error"`
}

func Classify(w http.ResponseWriter, r *http.Request) {
	numberParam := r.URL.Query().Get("number")
	if numberParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{
			Number: "",
			Error:  true,
		})
		return
	}

	num, err := strconv.Atoi(numberParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(APIError{
			Number: numberParam,
			Error:  true,
		})
		return
	}

	properties := []string{}
	if mathutils.IsArmStrong(num) {
		properties = append(properties, "armstrong")
	}
	if num%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}
	funFactText, err := funfact.GetFunFact(num)
	if err != nil || funFactText == "" {
		funFactText = fmt.Sprintf("%d is a number for which we're missing a fact (submit one to numbersapi at google mail!).", num)
	}

	res := ApiResponse{
		Number:     num,
		IsPrime:    mathutils.IsPrime(num),
		IsPerfect:  mathutils.IsPerfect(num),
		Properties: properties,
		DigitSum:   mathutils.DigitSum(num),
		FunFact:    funFactText,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
	}

}
