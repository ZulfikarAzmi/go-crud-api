# Go CRUD API

Go CRUD API adalah sebuah project sederhana yang menggunakan bahasa pemrograman Go untuk menyediakan layanan CRUD (Create, Read, Update, Delete) pada data film. API ini saya buat untuk memahami cara membuat RESTful API menggunakan framework seperti Gorilla Mux.

---

## **Fitur Utama:**
1. **Create Movie** - Menambahkan data film baru.
2. **Read Movies** - Mendapatkan daftar semua film.
3. **Read Movie** - Mendapatkan data film berdasarkan ID.
4. **Update Movie** - Mengubah data film berdasarkan ID.
5. **Delete Movie** - Menghapus data film berdasarkan ID.

---

## **Teknologi yang Digunakan:**
- **Bahasa Pemrograman:** Go
- **Framework Router:** Gorilla Mux

---

## **Struktur Data Movie:**
```json
{
  "id": "string",
  "isbn": "string",
  "title": "string",
  "director": {
    "firstname": "string",
    "lastname": "string"
  }
}
```

---

## **Endpoint API:**

| HTTP Method | Endpoint          | Deskripsi                |
|-------------|-------------------|--------------------------|
| GET         | `/movies`         | Mendapatkan semua film   |
| GET         | `/movies/{id}`    | Mendapatkan film spesifik|
| POST        | `/movies`         | Menambahkan film baru    |
| PUT         | `/movies/{id}`    | Memperbarui data film    |
| DELETE      | `/movies/{id}`    | Menghapus data film      |

---

## **Cara Menjalankan Project:**

1. **Clone Repository:**
   ```bash
   git clone https://github.com/ZulfikarAzmi/go-crud-api.git
   cd go-crud-api
   ```

2. **Jalankan Program:**
   ```bash
   go run main.go
   ```

3. **Akses API:**
   - Gunakan tool seperti Postman atau cURL untuk mengakses endpoint.
   - Default port: `http://localhost:8000`.

---

## **Contoh Request:**
### **POST /movies**
Request Body:
```json
{
  "isbn": "12345",
  "title": "My New Movie",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```
Response:
```json
{
  "id": "1",
  "isbn": "12345",
  "title": "My New Movie",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

---

## **Catatan:**
- Pastikan Anda sudah menginstal Go di komputer Anda.
- Jangan lupa untuk menambahkan file `.gitignore` untuk mengecualikan file yang tidak perlu di repository.

---

**Kontributor:**
Zulfikar Azmi

