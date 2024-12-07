package blockchain

    import (
      "encoding/json"
      "fmt"
      "net/http"
      "github.com/ethereum/go-ethereum/ethclient"
    )

    var cfg *config.Config

    func SetConfig(config *config.Config) {
      cfg = config
    }

    func InfoHandler(w http.ResponseWriter, r *http.Request) {
      client, err := ethclient.Dial(cfg.InfuraProjectID)
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      blockNumber, err := client.BlockNumber(r.Context())
      if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
      }

      response := map[string]interface{}{
        "blockNumber": blockNumber,
      }

      w.Header().Set("Content-Type", "application/json")
      json.NewEncoder(w).Encode(response)
    }
