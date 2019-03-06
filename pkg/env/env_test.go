package env

import (
    "os"
    "testing"
)

func TestGet(t *testing.T) {

    if val := Get("PORT", "8000"); val != "8000" {
        t.Errorf("expected 8000, got %v",  val)
    }

    if err := os.Setenv("PORT", "10000"); err != nil{
        t.Errorf("error setting env variable: %v", err)
    }

    if val := Get("PORT", "8000"); val != "10000" {
        t.Errorf("expected 10000, got %v",  val)
    }
}
