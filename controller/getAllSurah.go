package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Surah struct {
	ID         int    `json:"id"`
	SurahName  string `json:"surat_name"`
	SurahText  string `json:"surat_text"`
	Terjemahan string `json:"surat_terjemahan"`
	JumlahAyat int    `json:"count_ayat"`
}

func GetAllSurah(c *gin.Context) {
	file, err := os.Open("daftar-surah.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open JSON file"})
		return
	}
	defer file.Close()

	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read JSON file"})
		return
	}

	var response struct {
		Msg  string  `json:"msg"`
		Data []Surah `json:"data"`
	}

	if err := json.Unmarshal(byteValue, &response); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid JSON format"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": response.Msg,
		"data":    response.Data,
	})
}
