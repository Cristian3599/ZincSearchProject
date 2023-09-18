package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"enronmailbackend/models"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

const (
	ZINC_HOST = "http://localhost:4080"
	ZINC_USER = "admin"
	ZINC_PASS = "21jejAlo"
	INDEX     = "enron_mail"
)

func main() {
	// Crea un router para tu aplicación
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/emails", func(w http.ResponseWriter, r *http.Request) {
		queryResponse, err := getAllEmails()
		if err != nil {
			render.JSON(w, r, http.StatusInternalServerError)
		}
		render.JSON(w, r, queryResponse)
	})
	router.Get("/emails/{query}", func(w http.ResponseWriter, r *http.Request) {
		query := chi.URLParam(r, "query")
		queryResponse, err := getEmailsByQuery(query)
		if err != nil {
			render.JSON(w, r, http.StatusInternalServerError)
		}
		render.JSON(w, r, queryResponse)
	})
	//router.Get("/email/{id}", getEmail)

	// Escucha las conexiones en el puerto 3000
	log.Println("listening on port 3000")
	http.ListenAndServe(":3000", router)
}

func getAllEmails() (*models.QuerySearchResponse, error) {
	queryResponse := &models.QuerySearchResponse{}
	auth := ZINC_USER + ":" + ZINC_PASS
	bas64encoded_creds := base64.StdEncoding.EncodeToString([]byte(auth))
	//zinc_host := "https://playground.dev.zincsearch.com"
	zinc_url := ZINC_HOST + "/api/" + INDEX + "/_search"
	query := `{
        	"max_results": 20
		}`
	req, err := http.NewRequest("POST", zinc_url, strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}

	// Set headers
	//req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+bas64encoded_creds)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &queryResponse); err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error")
	}
	return queryResponse, nil
}

func getEmailsByQuery(query string) (*models.QuerySearchResponse, error) {
	queryResponse := &models.QuerySearchResponse{}
	auth := ZINC_USER + ":" + ZINC_PASS
	bas64encoded_creds := base64.StdEncoding.EncodeToString([]byte(auth))
	//zinc_host := "https://playground.dev.zincsearch.com"
	zinc_url := ZINC_HOST + "/api/" + INDEX + "/_search"
	var jsonStr []byte
	if query != "" {
		jsonStr = []byte(fmt.Sprintf(`{
			"search_type": "match",
			"query": {
					"term": "%v",
					"field": "body"
			},
			"from": 0,
			"max_results": 20,
			"_source": []
		}`, query))
	} else {
		jsonStr = []byte(`{ "max_results": 20 }`)
	}
	req, err := http.NewRequest("POST", zinc_url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal(err)
	}

	// Set headers
	//req.SetBasicAuth(user, password)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic "+bas64encoded_creds)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &queryResponse); err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Error")
	}
	return queryResponse, nil
}
