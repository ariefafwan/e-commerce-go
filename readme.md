# 🛒 E-Commerce-Go

E-Commerce-Go adalah aplikasi backend untuk platform e-commerce yang dibangun menggunakan bahasa pemrograman **Golang**. Aplikasi ini mencakup fitur seperti autentikasi pengguna, manajemen toko (up to multimerchant),manajemen produk & pelanggan, pemrosesan pesanan, serta integrasi beberapa sumber eksternal seperti cloudinary, raja ongkir, juga midtrans payment gateway. Proyek ini dibuat sebagai pembelajaran dan mudah dikembangkan lebih lanjut.

---

## 🚀 Tech Stack

### 🏢 Postgree

### 🎨 Backend: Golang

- Gin Framework
- GORM
- Cloudinary File System
- Validator
- JWT Auth (Access & Refresh Token) & Role Based Auth
- Middleware Token Validation
- Filtering and Pagination
- Raja Ongkir Integration (Shipping Cost)
- Midtrans Payment Gateway

## 📁 Structur Project

```bash
.
├── external/ # external api
├── internal/
│ └── controllers/
│ └── dto/ # data response
│ └── helpers/ # reusable function
│ └── models/ # modeling
│ └── repositories/ # bisnis logic
│ └── request/ # request logic
├── middleware/ # # JWT Auth, logging
├── pkg/ # reusable package aplikasi
├── routes/
│ └── api.go # API Routes
├── scraping/ # Modul scraping data eksternal (raja ongkir)
│ └── main.go
├── seeders/ # Data awal (seeding)
│ └── main.go
├── main.go # Entry point aplikasi
├── .env.example # Contoh file konfigurasi
├── .gitignore
├── README.md
```

# 🔧 Installasi & Development

1. Clone repository

```bash
   git clone https://github.com/ariefafwan/e-commerce-go.git
   cd e-commerce.go
   cp .env.example .env
```

Edit file .env sesuai konfigurasi lokal kamu (DATABASE_URL, PORT, dll.) juga sesuaikan environment lainnya

2. Setup Dependencies

```bash
   go mod tidy

   # scraping data
   cd scraping
   cp ~/e-commerce-go/.env .env
   go run main.go --scraping=all
   cd ..

   # seed data
   cd seeders
   cp ~/e-commerce-go/.env .env
   go run main.go --seed=all
   cd ..

   # build or run
   go run main.go # run dev
   go build -o {nama-app} main.go # build
```

Go akan berjalan di port sesuai env anda

3. Example Endpoints

| Method | Endpoint                     | Deskripsi                             |
| ------ | ---------------------------- | ------------------------------------- |
| POST   | /api/register                | Registrasi akun baru pengguna         |
| POST   | /api/login                   | Login dan mendapatkan JWT             |
| POST   | /api/refresh                 | Refresh Token Anda                    |
| POST   | /api/logout                  | Logout dan revoke semua token user    |
| GET    | /api/master-pelanggan/       | Mendapatkan semua pelanggan           |
| GET    | /api/master-kategori-produk/ | Mendapatkan semua produk              |
| POST   | /api/master-produk           | Menambah produk baru                  |
| POST   | /api/keranjang               | Menambahkan item ke dalam keranjang   |
| POST   | /api/transaksi/kalkulasi/    | Menghitung total yang akan di chekout |
| DELETE | /api/transaksi/              | Membuat transaksi baru dari kalkulasi |

Dokumentasi Endpoint ada di : [POSTMAN](https://documenter.getpostman.com/view/26198524/2sB3B7MYf2)

4. Login

For Admin, login with:

- email: admin@admin.com
- password: 123
  (akun ini bisa anda ubah saat seed awal)

For Default Pelanggan, login with:

- email: pelanggan2@gmail.com
- password: 123
  (akun ini bisa anda ubah saat seed awal)
  (kamu juga bisa register untuk pelanggan baru)

# 🧩 Relasi Utama

```bash
.
User
└── Personal Acces Token (sudah ada set time revoke)
└── Pelanggan (Roles Pelanggan)
Toko -> Untuk mengatur alamat toko, dan aturan lainnya seperti pajak
Pelanggan
└── AlamatPengiriman (banyak, satu is_default = true)
Kategori Produk
├── id_parent (self relationship)
└── DataProduk (many to many)
Produk
├── GaleriGambar (many)
├── Variant (many, bisa berupa kombinasi ukuran/warna)
└── DataKategoriProduk (many to many)
Keranjang
└── Item: berdasarkan ProdukVariant + kuantitas
Transaksi
└── ItemTransaksi: copy dari item keranjang terpilih
└── Status: Pending, Cancelled, Expired, Paid, Complete
```

# 👥 Kontribusi

- Fork repository ini
- Buat branch fitur baru: git checkout -b fitur-anda
- Commit perubahan: git commit -m 'Menambahkan fitur xyz'
- Push ke branch: git push origin fitur-anda
- Buat Pull Request

# 📄 Lisensi

Proyek ini open source. Lihat file [LICENSE](https://choosealicense.com/licenses/mit/) untuk informasi lebih lanjut.

# 👤 Author

## Teuku M Arief Afwan

### 📄 GitHub: [@ariefafwan](https://github.com/ariefafwan/)

### 👋 Instagram: [@teukuafwan](https://www.instagram.com/teukuafwan/)

### 🔗 Personal Website: [teukuafwan](teukuafwan.my.id)

Repository ini juga live di [e-commerce-app](https://ws-ecommerce-go.teukuafwan.my.id/)
