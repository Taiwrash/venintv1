package main_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	main "github.com/Taiwrash/venintv1"
)

func TestCreateMessageHandler(t *testing.T) {
	// Prepare the request body
	message := main.Message{Content: "Hello, world!"}
	requestBody, _ := json.Marshal(message)
	request := httptest.NewRequest("POST", "/messages", bytes.NewBuffer(requestBody))
	responseRecorder := httptest.NewRecorder()

	// Invoke the handler
	main.CreateMessageHandler(responseRecorder, request)

	// Verify the response
	assert.Equal(t, http.StatusCreated, responseRecorder.Code)

	// Parse and validate the response body
	var responseMessage main.Message
	json.Unmarshal(responseRecorder.Body.Bytes(), &responseMessage)
	assert.Equal(t, message.Content, responseMessage.Content)
}

func TestGetMessageHandler(t *testing.T) {
	// Prepare the request
	request := httptest.NewRequest("GET", "/messages/{id}", nil)
	request = mux.SetURLVars(request, map[string]string{"id": "123"})
	responseRecorder := httptest.NewRecorder()

	// Invoke the handler
	main.GetMessageHandler(responseRecorder, request)

	// Verify the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Parse and validate the response body
	var message main.Message
	json.Unmarshal(responseRecorder.Body.Bytes(), &message)
	assert.Equal(t, "Hello, world!", message.Content)
}

func TestGetAccountStateHandler(t *testing.T) {
	// Prepare the request
	request := httptest.NewRequest("GET", "/accounts/{address}", nil)
	request = mux.SetURLVars(request, map[string]string{"address": "0x123abc"})
	responseRecorder := httptest.NewRecorder()

	// Invoke the handler
	main.GetAccountStateHandler(responseRecorder, request)

	// Verify the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Parse and validate the response body
	var accountState main.AccountState
	json.Unmarshal(responseRecorder.Body.Bytes(), &accountState)
	assert.Equal(t, "0x123abc", accountState.Address)
	assert.Equal(t, 100, accountState.Balance)
}

func TestQueryBlockchainDataHandler(t *testing.T) {
	// Prepare the request
	request := httptest.NewRequest("GET", "/blockchain", nil)
	responseRecorder := httptest.NewRecorder()

	// Invoke the handler
	main.QueryBlockchainDataHandler(responseRecorder, request)

	// Verify the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Parse and validate the response body
	var blockchainData main.BlockchainData
	json.Unmarshal(responseRecorder.Body.Bytes(), &blockchainData)
	assert.Len(t, blockchainData.Blocks, 2)
	assert.Len(t, blockchainData.Transactions, 2)
	assert.Len(t, blockchainData.Messages, 2)
}

func TestSubscribeToEventsHandler(t *testing.T) {
	// Prepare the request
	request := httptest.NewRequest("POST", "/subscribe", nil)
	responseRecorder := httptest.NewRecorder()

	// Invoke the handler
	main.SubscribeToEventsHandler(responseRecorder, request)

	// Verify the response
	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	// Parse and validate the response body
	var response map[string]string
	json.Unmarshal(responseRecorder.Body.Bytes(), &response)
	assert.Equal(t, "Subscribed to events", response["message"])
}

// Implement similar tests for the remaining handlers

