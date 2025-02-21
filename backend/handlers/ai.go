package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RecommendTask(c *gin.Context) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	requestData := map[string]string{
		"prompt": "Generate task ideas based on project requirements",
	}

	body, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI request failed"})
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	c.JSON(http.StatusOK, result)
}
