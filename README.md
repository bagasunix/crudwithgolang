# crudwithgolang

# Requirements / Pre-requisites

- Install semua dependencies: `go get .`
- Jika ingin menjalankan REST API ini secara lokal, pastikan ada:
  - Instalasi Go di komputer
  - Service MongoDB dan Minio yang berjalan
- Buat environment variable dalam `.env` file (check `.env.example` untuk variable apa saja yang dibutuhkan, sesuaikan dengan kebutuhan)

# Run the server

`go run .`

Command diatas akan menjalankan REST API di port :8080 (bisa di set lewat env variable `HOST_PORT`)
