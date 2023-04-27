package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/hgfinance/models"
	"github.com/joho/godotenv"
)

func Dollar(c *gin.Context) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Erro ao carregar variavel de ambiente para carregar URL", err)
		return
	}

	URL := os.Getenv("URL")
	url, err := http.Get(URL)
	if err != nil {
		log.Println("Erro ao ler URL", err, url.Body)
		return
	}

	data, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Println("Erro ao ler body", err)
		return
	}

	var result models.USD
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println("Erro ao realizar unmarshal do data", err)
		return
	}
	json.MarshalIndent(result, "", "    ")

	fmt.Println("Preço:", result.USDBRL.Bid)
	fmt.Println("Moeda:", result.USDBRL.Name)
	fmt.Println("Data:", result.USDBRL.Date)

	c.JSON(http.StatusOK, gin.H{
		"Data":  result.USDBRL.Date,
		"Moeda": result.USDBRL.Name,
		"Preço": result.USDBRL.Bid,
	})
}

func BitCoin(c *gin.Context) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("Erro ao carregar variavel de ambiente para carregar URL", err)
		return
	}

	URL := os.Getenv("URLBIT")
	url, err := http.Get(URL)
	if err != nil {
		log.Println("Erro ao ler URL", err, url.Body)
		return
	}

	data, err := ioutil.ReadAll(url.Body)
	if err != nil {
		log.Println("Erro ao ler body", err)
		return
	}

	var result models.BitCoin
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Println("Erro ao realizar unmarshal do data", err)
		return
	}
	json.MarshalIndent(result, "", "    ")

	fmt.Println("Preço:", result.BTCBRL.Bid)
	fmt.Println("Moeda:", result.BTCBRL.Name)
	fmt.Println("Data:", result.BTCBRL.Date)

	c.JSON(http.StatusOK, gin.H{
		"Data":  result.BTCBRL.Date,
		"Moeda": result.BTCBRL.Name,
		"Preço": result.BTCBRL.Bid,
	})
}
