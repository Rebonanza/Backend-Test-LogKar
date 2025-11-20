package main

import (
	"log"

	"github.com/local/be-test-logkar/internal/config"
	"github.com/local/be-test-logkar/internal/db"
	"github.com/local/be-test-logkar/internal/user"
)

func main() {
	cfg := config.NewConfig()
	dbConn, err := db.NewGormDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}

	samples := []user.User{
		{Email: "user1@mail.com", Name: "user1"},
	}

	for _, s := range samples {
		if err := dbConn.Create(&s).Error; err != nil {
			log.Printf("seed insert failed for %s: %v", s.Email, err)
		} else {
			log.Printf("inserted seed user %s", s.Email)
		}
	}

	products := []struct {
		Name     string
		Type     string
		Flavor   string
		Size     string
		Price    int
		Quantity int
	}{
		{Name: "Keripik Pangsit", Type: "Keripik Pangsit", Flavor: "Jagung Bakar", Size: "Small", Price: 10000, Quantity: 50},
		{Name: "Keripik Pangsit", Type: "Keripik Pangsit", Flavor: "Original", Size: "Medium", Price: 25000, Quantity: 40},
		{Name: "Keripik Pangsit", Type: "Keripik Pangsit", Flavor: "Keju", Size: "Large", Price: 35000, Quantity: 30},
	}

	for _, p := range products {
		prod := struct {
			Name     string
			Type     string
			Flavor   string
			Size     string
			Price    int
			Quantity int
		}{p.Name, p.Type, p.Flavor, p.Size, p.Price, p.Quantity}
		if err := dbConn.Exec(`INSERT INTO products (name, type, flavor, size, price, quantity) VALUES (?, ?, ?, ?, ?, ?) ON CONFLICT DO NOTHING`, prod.Name, prod.Type, prod.Flavor, prod.Size, prod.Price, prod.Quantity).Error; err != nil {
			log.Printf("failed inserting product %s: %v", prod.Name, err)
		}
	}

	custs := []struct{ Name string }{{"Customer A"}, {"Customer B"}}
	for _, c := range custs {
		if err := dbConn.Exec(`INSERT INTO customers (name, points) VALUES (?, 0) ON CONFLICT DO NOTHING`, c.Name).Error; err != nil {
			log.Printf("failed inserting customer %s: %v", c.Name, err)
		}
	}
	log.Println("seeding complete")
}
