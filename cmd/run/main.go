package main

import (
	"encoding/json"
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"io"
	"log"
	"net/http"
)

type Subscription struct {
	Endpoint string            `json:"endpoint"`
	Keys     map[string]string `json:"keys"`
}

var subscriptions []Subscription

func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	var subscription Subscription

	if err := json.NewDecoder(r.Body).Decode(&subscription); err != nil {
		fmt.Printf("Error: %s", err)
		http.Error(w, "Невалидный формат подписки", http.StatusBadRequest)
		return
	}

	fmt.Printf(subscription.Endpoint)

	subscriptions = append(subscriptions, subscription)
	w.WriteHeader(http.StatusOK)
}

func Notify() {
	fmt.Println(len(subscriptions))
	for _, sub := range subscriptions {
		s := &webpush.Subscription{
			Endpoint: sub.Endpoint,
			Keys: webpush.Keys{
				Auth:   sub.Keys["auth"],
				P256dh: sub.Keys["p256dh"],
			},
		}

		resp, err := webpush.SendNotification([]byte(fmt.Sprintf("your key is %s", s.Keys.Auth)), s, &webpush.Options{
			Subscriber:      "leonov.sas2018@yandex.ru",
			VAPIDPublicKey:  "BG5LblQ_TNwE5hegYZVWaBN45TegcepZUB97Md0x-BGYJxkX5neXwP-Ihcc1pjBw7SzEvOC_ZSQzBfIhw2daEzg",
			VAPIDPrivateKey: "-yfFUsO5Bww06PwDnxODvvAnuWejsx5qFt4f2adVUas",
			TTL:             60,
		})

		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		str, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		fmt.Printf("Response: %s\n", str)
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))

	r.Post("/subscribe", subscribeHandler)

	go func() {
		for {
			var s string
			fmt.Scan(&s)
			Notify()
		}
	}()

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
