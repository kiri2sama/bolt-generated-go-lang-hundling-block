package main

    import (
      "log"
      "net/http"
      "blockchain-backend/internal/blockchain"
      "blockchain-backend/internal/smartcontract"
      "blockchain-backend/pkg/config"
    )

    func main() {
      // Load configuration
      cfg, err := config.LoadConfig()
      if err != nil {
        log.Fatalf("Failed to load configuration: %s\n", err.Error())
      }

      // Pass the configuration to the blockchain and smartcontract modules
      blockchain.SetConfig(cfg)
      smartcontract.SetConfig(cfg)

      http.HandleFunc("/blockchain/info", blockchain.InfoHandler)
      http.HandleFunc("/smartcontract/interact", smartcontract.InteractHandler)

      log.Println("Starting server on :8080")
      if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
      }
    }
