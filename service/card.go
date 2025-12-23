package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"template/config"
	"template/model"
	"template/repository"
	"time"
)

type CardService interface {
	ScrapCards() error
}

type cardService struct {
	cardRepo repository.CardRepository
}

func NewCardService(cardRepo repository.CardRepository) CardService {
	return &cardService{cardRepo}
}

type PokemonTCGResponse struct {
	Data       []PokemonCard `json:"data"`
	Page       int           `json:"page"`
	PageSize   int           `json:"pageSize"`
	Count      int           `json:"count"`
	TotalCount int           `json:"totalCount"`
}

type PokemonCard struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Supertype   string   `json:"supertype"`
	Subtypes    []string `json:"subtypes"`
	HP          string   `json:"hp"`
	Types       []string `json:"types"`
	EvolvesFrom string   `json:"evolvesFrom"`
	Images      struct {
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"images"`
	Rarity string `json:"rarity"`
	Artist string `json:"artist"`
	Set    struct {
		ID string `json:"id"`
	} `json:"set"`
	Number string `json:"number"`
}

func (s *cardService) ScrapCards() error {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	page := 1
	pageSize := 1 // Max page size allowed by API

	for {
		url := fmt.Sprintf("https://api.pokemontcg.io/v2/cards?page=%d&pageSize=%d", page, pageSize)
		fmt.Printf("Fetching page %d...\n", page)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}

		// It's good practice to add a User-Agent, though possibly optional here
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")

		if config.PokemonApiKey != "" {
			req.Header.Set("X-Api-Key", config.PokemonApiKey)
		}

		fmt.Println("req: ", req)

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		fmt.Println("resp: ", resp)

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("api returned status: %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		fmt.Println("body: ", string(body))

		var result PokemonTCGResponse
		if err := json.Unmarshal(body, &result); err != nil {
			return err
		}

		if len(result.Data) == 0 {
			break
		}

		for _, apiCard := range result.Data {
			subtypesJson, _ := json.Marshal(apiCard.Subtypes)
			typesJson, _ := json.Marshal(apiCard.Types)
			imagesJson, _ := json.Marshal(apiCard.Images)

			card := &model.Card{
				ID:          apiCard.ID,
				Name:        apiCard.Name,
				Supertype:   apiCard.Supertype,
				Subtypes:    subtypesJson,
				HP:          apiCard.HP,
				Types:       typesJson,
				EvolvesFrom: apiCard.EvolvesFrom,
				Images:      imagesJson,
				Rarity:      apiCard.Rarity,
				Artist:      apiCard.Artist,
				SetID:       apiCard.Set.ID,
				Number:      apiCard.Number,
			}

			// Save to DB
			if err := s.cardRepo.Create(card); err != nil {
				// Log error but continue? Or fail?
				// For now let's just log and continue
				fmt.Printf("Failed to save card %s: %v\n", card.ID, err)
			}
		}

		if result.Count < pageSize {
			break
		}

		// Create a small delay to avoid hitting rate limits hard if running without key
		time.Sleep(100 * time.Millisecond)

		page++
	}

	return nil
}
