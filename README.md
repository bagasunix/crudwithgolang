# CRUD WIth Golang

# Requirements / Pre-requisites

- Install semua dependencies: `go get .`
- Jika ingin menjalankan REST API ini secara lokal, pastikan ada:
  - Instalasi Go di komputer
  - Service MongoDB dan Minio yang berjalan
- Buat environment variable dalam `.env` file (check `.env.example` untuk variable apa saja yang dibutuhkan, sesuaikan dengan kebutuhan)

# Run the server

`go run .`

Command diatas akan menjalankan REST API di port :8080 (bisa di set lewat env variable `HOST_PORT`)

# List of API Routers

| Route | HTTP   | Description                       |
| ----- | ------ | --------------------------------- |
| /     | POST   | Route used to create a new item   |
| /:id  | GET    | Route used to get item all        |
| /:id  | GET    | Route used to get item by slug    |
| /:id  | PATCH  | Route used to update item by slug |
| /:id  | DELETE | Route used to delete item by slug |

### Usage

---

POST

`/api/user/register `

- **Body**

```
{
    "name":         "Laptop Asus VivoBook 14",
		"price":        100000000,
		"description":  "laptop terbaru",
    "type":         "Laptop"
}
```

GET ALL ITEM

`/`

GET BY SLUG

`/:id`

- **Params**

```
/asus-vivobook-14
```

PATCH BY SLUG

`/:id`

- **Params**

```
/asus-vivobook-14
```

- **Body**

```
{
    "name":         "Laptop Asus VivoBook 15",
		"price":        100000000,
		"description":  "laptop terbaru",
    "type":         "Laptop"
}
```

DELETE BY SLUG

`/:id`

- **Params**

```
/asus-vivobook-14
```
