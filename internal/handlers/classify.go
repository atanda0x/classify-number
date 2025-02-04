package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	Number  string `json:"number"`
	Error   bool   `json:"error"`
	Message string `json:"message,omitempty"`
}


func Classify(w http.ResponseWriter, r *http.Request) {
	numberParam := r.URL.Query().Get("number")
	if numberParam == "" {
		http.Error(w, `{"number": "", "error": true, "message": "Missing number parameter"}`, http.StatusBadRequest)
		return
	}

	// Convert string to integer
	num, err := strconv.Atoi(numberParam)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"number": "%s", "error": true, "message": "Invalid number format"}`, numberParam), http.StatusBadRequest)
		return
	}

	
	properties := []string{}
	if mathutils.IsArmstrong(num) {
		properties = append(properties, "armstrong")
	}
	if num%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	funFactText := generateFunFact(num)

	
	response := ApiResponse{
		Number:     num,
		IsPrime:    mathutils.IsPrime(num),
		IsPerfect:  mathutils.IsPerfect(num),
		Properties: properties,
		DigitSum:   mathutils.DigitSum(num),
		FunFact:    funFactText,
	}

	
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, `{"error": true, "message": "Internal Server Error"}`, http.StatusInternalServerError)
	}
}

func generateFunFact(num int) string {
	if mathutils.IsArmstrong(num) {
		digits := strconv.Itoa(num)
		sumParts := ""
		sum := 0
		for _, digit := range digits {
			d := int(digit - '0')
			sum += d * d * d
			sumParts += fmt.Sprintf("%d^3 + ", d)
		}
		sumParts = sumParts[:len(sumParts)-3] // Remove trailing " + "
		return fmt.Sprintf("%d is an Armstrong number because %s = %d.", num, sumParts, sum)
	}
	if mathutils.IsPrime(num) {
		return fmt.Sprintf("%d is a prime number because it has exactly two factors: 1 and %d.", num, num)
	}
	if num%2 == 0 {
		return fmt.Sprintf("%d is an even number because it is divisible by 2.", num)
	}
	return fmt.Sprintf("%d is an odd number because it is not divisible by 2.", num)
}
