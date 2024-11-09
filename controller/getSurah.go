package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSurah(c *gin.Context) {
	// Get the Surah ID from the URL parameter
	surahIDStr := c.Param("id")
	surahID, err := strconv.Atoi(surahIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Surah ID"})
		return
	}

	// Define the structures for Surah and SurahMeta
	type Surah struct {
		AyaId              int    `json:"aya_id"`
		AyaNumber          int    `json:"aya_number"`
		AyaText            string `json:"aya_text"`
		SuraId             int    `json:"sura_id"`
		JuzId              int    `json:"juz_id"`
		PageNumber         int    `json:"page_number"`
		TranslationAyaText string `json:"translation_aya_text"`
	}

	type SurahMeta struct {
		ID         int    `json:"id"`
		SurahName  string `json:"surat_name"`
		SurahText  string `json:"surat_text"`
		Terjemahan string `json:"surat_terjemahan"`
		JumlahAyat int    `json:"count_ayat"`
	}

	// Load the specific Surah file (e.g., surah/1.json)
	surahFilePath := fmt.Sprintf("./surah/%d.json", surahID)
	surahFile, err := os.Open(surahFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not open Surah file: %s", surahFilePath)})
		return
	}
	defer surahFile.Close()

	// Read the content of the Surah file
	surahFileContent, err := os.ReadFile(surahFilePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read Surah file"})
		return
	}

	// Load the daftar-surah.json (contains the metadata)
	metaFile, err := os.Open("daftar-surah.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not open daftar-surah.json"})
		return
	}
	defer metaFile.Close()

	metaFileContent, err := os.ReadFile("daftar-surah.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not read daftar-surah.json"})
		return
	}

	// Define response struct for Surah and its metadata
	var surahResponse struct {
		Msg  string  `json:"msg"`
		Data []Surah `json:"data"`
	}

	var metaResponse struct {
		Msg  string      `json:"msg"`
		Data []SurahMeta `json:"data"`
	}

	// Unmarshal Surah data
	if err := json.Unmarshal(surahFileContent, &surahResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Surah file format"})
		return
	}

	// Unmarshal Surah metadata (from daftar-surah.json)
	if err := json.Unmarshal(metaFileContent, &metaResponse); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid daftar-surah.json format"})
		return
	}

	// Find the matching metadata for the Surah ID (if applicable)
	var surahMeta SurahMeta
	for _, meta := range metaResponse.Data {
		if meta.ID == surahID {
			surahMeta = meta
			break
		}
	}

	// Return the Surah data and its metadata
	c.JSON(http.StatusOK, gin.H{
		"message": surahResponse.Msg,
		"data":    surahResponse.Data,
		"meta":    surahMeta,
	})
}
