package main

import (
	"e-commerce-go/pkg"
	"e-commerce-go/scraping/scraping"
	"flag"
	"fmt"
)

func main() {
	pkg.ConnectDB()
	
	scrapingProvince := flag.String("scraping", "", "scraping_from_raja_ongkir")
	flag.Parse()
	switch *scrapingProvince {
		case "provinsi":
			fmt.Println("Scraping Master Provinsi...")
			scraping.ScrapingProvinsi(pkg.DB)
		case "kota":
			fmt.Println("Scraping Master Kota...")
			scraping.ScrapingKota(pkg.DB)
		case "all":
			fmt.Println("Scraping Master Provinsi & Kota...")
			scraping.ScrapingProvinsi(pkg.DB)
			scraping.ScrapingKota(pkg.DB)
		default:
			break
	}
}