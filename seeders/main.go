package main

import (
	"e-commerce-go/pkg"
	"e-commerce-go/seeders/seeders"
	"flag"
)

func main() {
	pkg.ConnectDB()
	
	seedingFlag := flag.String("seed", "all", "seeding_table")
	flag.Parse()
	switch *seedingFlag {
	case "master_toko":
		seeders.SeedMasterToko(pkg.DB)
	case "user":
		seeders.SeedUser(pkg.DB)
	case "admin":
		seeders.SeedAdmin(pkg.DB)
	case "all":
		seeders.SeedMasterToko(pkg.DB)
		seeders.SeedUser(pkg.DB)
		seeders.SeedAdmin(pkg.DB)
	default:
		break
	}
}