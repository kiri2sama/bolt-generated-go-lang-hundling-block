package config

    import (
      "fmt"
      "os"
      "github.com/joho/godotenv"
    )

    type Config struct {
      InfuraProjectID string
      ContractAddress string
    }

    func LoadConfig() (*Config, error) {
      // Load environment variables from .env file
      err := godotenv.Load()
      if err != nil {
        return nil, fmt.Errorf("failed to load .env file: %s", err.Error())
      }

      config := &Config{
        InfuraProjectID: os.Getenv("INFURA_PROJECT_ID"),
        ContractAddress: os.Getenv("CONTRACT_ADDRESS"),
      }

      if config.InfuraProjectID == "" {
        return nil, fmt.Errorf("INFURA_PROJECT_ID environment variable not set")
      }

      if config.ContractAddress == "" {
        return nil, fmt.Errorf("CONTRACT_ADDRESS environment variable not set")
      }

      return config, nil
    }
