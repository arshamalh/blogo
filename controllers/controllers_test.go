package controllers

import (
	"testing"
	"go.uber.org/zap"
)

func TestBasicAttributesInitialization(t *testing.T) {
	// Mock database and logger for testing purposes
	mockDB := 
	mockLogger, _ := zap.NewDevelopment()
	basicAttrs := &basicAttributes{
		db:     mockDB,
		logger: mockLogger,
		Gl:     mockLogger, 
	}
	
	if basicAttrs.db == nil {
		t.Error("Expected non-nil database, got nil")
	}

	if basicAttrs.logger == nil {
		t.Error("Expected non-nil logger, got nil")
	}

	if basicAttrs.Gl == nil {
		t.Error("Expected non-nil Gl logger, got nil")
	}
}


