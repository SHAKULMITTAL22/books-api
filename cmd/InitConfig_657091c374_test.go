package main

import (
    "booksApi"
    "booksApi/pkg/handler"
    "booksApi/pkg/repository"
    "booksApi/pkg/service"
    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
    "github.com/spf13/viper"
    "log"
    "os"
    "testing"
)

func TestInitConfig_657091c374(t *testing.T) {
    // Test case 1: Success case
    err := initConfig()
    if err != nil {
        t.Errorf("initConfig() failed: %v", err)
    } else {
        t.Log("initConfig() succeeded")
    }

    // Test case 2: Failure case
    viper.Reset()
    err = initConfig()
    if err == nil {
        t.Errorf("initConfig() succeeded unexpectedly")
    } else {
        t.Logf("initConfig() failed as expected: %v", err)
    }
}
