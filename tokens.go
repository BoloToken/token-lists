package tokens

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://api.uniswap.com/tokens")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var uniswapTokens []Token
	err = json.Unmarshal(data, &uniswapTokens)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	localData, err := os.ReadFile("localTokens.json")
	if err != nil {
		fmt.Println("Error reading localTokens.json file:", err)
		return
	}

	var localTokens []Token
	err = json.Unmarshal(localData, &localTokens)
	if err != nil {
		fmt.Println("Error decoding local JSON:", err)
		return
	}

	updatedLocalData, err := json.Marshal(localTokens)
	if err != nil {
		fmt.Println("Error encoding updated JSON:", err)
		return
	}

	err = os.WriteFile("localTokens.json", updatedLocalData, 0644)
	if err != nil {
		fmt.Println("Error writing file localTokens.json:", err)
		return
	}
}

type Token struct {
	Name    string `json:"name"`
	Symbol  string `json:"symbol"`
	Address string `json:"address"`
}
