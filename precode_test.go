package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerCorrectResponse(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=4", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.NotEmpty(t, responseRecorder.Body)

}

func TestMainHandlerIncorrectCity(t *testing.T) {
	expectedBody := "wrong city value"
	req := httptest.NewRequest("GET", "/cafe?city=saint-peterburg&count=4", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	serverBody := responseRecorder.Body.String()
	assert.Equal(t, expectedBody, serverBody)
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	correctResponse := []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"}
	req := httptest.NewRequest("GET", "/cafe?city=moscow&count=777", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	serverResponse := strings.Split(body, ",")

	assert.Equal(t, correctResponse, serverResponse)
}
