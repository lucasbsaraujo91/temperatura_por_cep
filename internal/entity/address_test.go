package entity

import (
	"testing"
)

func TestNewConsultZipCode_Success(t *testing.T) {
	address, err := NewConsultZipCode("12345678", "Rua Exemplo", "Bairro Exemplo", "São Paulo", "SP")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if address.ZipCode != "12345678" {
		t.Errorf("Expected ZipCode to be '12345678', got %v", address.ZipCode)
	}
	if address.Street != "Rua Exemplo" {
		t.Errorf("Expected Street to be 'Rua Exemplo', got %v", address.Street)
	}
	if address.Neighborhood != "Bairro Exemplo" {
		t.Errorf("Expected Neighborhood to be 'Bairro Exemplo', got %v", address.Neighborhood)
	}
	if address.City != "São Paulo" {
		t.Errorf("Expected City to be 'São Paulo', got %v", address.City)
	}
	if address.State != "SP" {
		t.Errorf("Expected State to be 'SP', got %v", address.State)
	}
}

func TestNewConsultZipCode_InvalidZipCode(t *testing.T) {
	_, err := NewConsultZipCode("12345", "Rua Exemplo", "Bairro Exemplo", "São Paulo", "SP")
	if err == nil {
		t.Fatal("Expected an error for invalid ZipCode, got none")
	}
}

func TestNewConsultZipCode_NonNumericZipCode(t *testing.T) {
	_, err := NewConsultZipCode("12345678a", "Rua Exemplo", "Bairro Exemplo", "São Paulo", "SP")
	if err == nil {
		t.Fatal("Expected an error for non-numeric ZipCode, got none")
	}
}

func TestAddress_IsValid_Success(t *testing.T) {
	address := &Address{ZipCode: "12345678"}
	err := address.IsValid()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestAddress_IsValid_InvalidLength(t *testing.T) {
	address := &Address{ZipCode: "123"}
	err := address.IsValid()
	if err == nil {
		t.Fatal("Expected an error for invalid ZipCode length, got none")
	}
}

func TestAddress_IsValid_NonNumeric(t *testing.T) {
	address := &Address{ZipCode: "1234abcd"}
	err := address.IsValid()
	if err == nil {
		t.Fatal("Expected an error for non-numeric ZipCode, got none")
	}
}
