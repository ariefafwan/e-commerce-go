package seeders

import (
	"e-commerce-go/pkg"
	"flag"
)

func main() {
	seedingFlag := flag.String("seed", "MasterToko", "seeding_master_toko")
	flag.Parse()
	if *seedingFlag == "MasterToko" {
		SeedMasterToko(pkg.DB)
	}
}