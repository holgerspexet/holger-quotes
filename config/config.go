package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// StorageTypeMemory is the storage type for memory storage
const StorageTypeMemory string = "memory"

// StorageTypeSQLight is the storage type for sqlite storage
const StorageTypeSQLight string = "sqlite"

// Config holds all configuration values for the application
type Config struct {
	Port        int
	StorageType string
	SQLightPath string
	Hosting     string
}

// LoadConfig creates a new Config object from available ENV variables
func LoadConfig() Config {
	storageType := loadStorageType()
	config := Config{
		Port:        loadPort(),
		StorageType: storageType,
		SQLightPath: loadSQLightPath(storageType),
		Hosting:     loadHosting(),
	}
	log.Printf("Config: %+v", config)
	return config
}

func loadPort() int {
	portString := os.Getenv("HOLGER_QUOTES_PORT")
	if portString == "" {
		portString = "9000"
		log.Printf("No port provided, defaulting to: %s", portString)
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		log.Fatal(err)
	}
	return port
}

func loadStorageType() string {
	storageType := os.Getenv("HOLGER_QUOTES_STORAGE_TYPE")
	if storageType == "" {
		storageType = StorageTypeSQLight
		log.Printf("No storageType provided, defaulting to: %s", storageType)
	}
	if storageType != StorageTypeSQLight && storageType != StorageTypeMemory {
		log.Fatalf("Invalid storageType: %s", storageType)
	}
	return storageType
}

func loadSQLightPath(storageType string) string {
	if storageType != StorageTypeSQLight {
		return ""
	}

	path := os.Getenv("HOLGER_QUOTES_SQLIGHT_PATH")
	if path == "" {
		path, _ = filepath.Abs("./sqlite3.sql")
		log.Printf("No sqlightPath provided, defaulting to: %s", path)
	}
	return path
}

func loadHosting() string {
	hosting := os.Getenv("HOLGER_QUOTES_HOSTING")
	if hosting == "" {
		hosting = "/"
		log.Printf("No hosting provided, defaulting to: %s", hosting)
	}
	if hosting != "/" {
		hosting = strings.ReplaceAll(hosting, "//", "/")
	}
	return hosting
}
