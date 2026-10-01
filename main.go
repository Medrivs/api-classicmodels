package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ProductCode string  `json:"productCode"`
	ProductName string  `json:"productName"`
	ProductLine string  `json:"productLine"`
	BuyPrice    float64 `json:"buyPrice"`
}

var db *sql.DB

func main() {
	var err error
	
	// AQUÍ CAMBIA "TU_CONTRASEÑA" POR TU CONTRASEÑA DE MYSQL
	dsn := "root:Alanmbritos1@tcp(127.0.0.1:3306)/classicmodels"
	
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error abriendo la base de datos:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error conectando a la base de datos. ¿Pusiste bien la contraseña?:", err)
	}

	http.HandleFunc("/api/products", getProductsHandler)

	fmt.Println("Servidor ejecutándose en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // Permite que se pueda leer desde un navegador fácilmente

	// Traemos 10 productos de la base de datos
	query := "SELECT productCode, productName, productLine, buyPrice FROM products LIMIT 10"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error en la consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []Product

	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.ProductCode, &p.ProductName, &p.ProductLine, &p.BuyPrice); err != nil {
			http.Error(w, "Error leyendo los datos", http.StatusInternalServerError)
			return
		}
		products = append(products, p)
	}

	json.NewEncoder(w).Encode(products)
}