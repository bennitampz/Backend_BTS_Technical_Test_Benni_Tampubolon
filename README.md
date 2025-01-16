# Backend_BTS_Technical_Test_Benni_Tampubolon

## Deskripsi

Repositori ini berisi implementasi backend untuk aplikasi checklist. Aplikasi ini menyediakan fitur untuk registrasi dan login pengguna, serta pengelolaan checklist dan item-item di dalamnya.

## Fitur

### Otentikasi
- **Registrasi Pengguna:** Membuat akun pengguna baru.
- **Login Pengguna:** Masuk ke aplikasi dengan akun yang sudah terdaftar.

### Checklist
- **Melihat Semua Checklist:** Mendapatkan daftar semua checklist milik pengguna yang sedang login.
- **Membuat Checklist Baru:** Membuat checklist baru.
- **Melihat Detail Checklist:** Mendapatkan detail sebuah checklist berdasarkan ID.
- **Menghapus Checklist:** Menghapus checklist berdasarkan ID.

### Item Checklist
- **Melihat Semua Item:** Mendapatkan daftar semua item dalam sebuah checklist.
- **Membuat Item Baru:** Membuat item baru dalam sebuah checklist.
- **Memperbarui Item:** Memperbarui detail sebuah item dalam checklist.
- **Menghapus Item:** Menghapus item dari sebuah checklist.

# Cara Menjalankan Aplikasi Backend Checklist

Dokumen ini menjelaskan langkah-langkah untuk menjalankan aplikasi backend checklist yang dibuat dengan Go.

## Prasyarat

Pastikan Anda telah menginstal Go (Golang) di sistem Anda. Anda dapat mengunduh dan menginstal Go dari [situs resmi Go](https://golang.org/dl/).

## Cara Mengunduh

1.  **Clone Repositori:**
    Anda dapat mengunduh repositori ini dengan menggunakan perintah `git clone` di terminal Anda:
    ```bash
    git clone https://github.com/bennitampz/Backend_BTS_Technical_Test_Benni_Tampubolon
    ```


2.  **Pindah ke Direktori Proyek:**
    Setelah berhasil mengunduh, pindah ke direktori proyek:
    ```bash
    cd Backend_BTS_Technical_Test_Benni_Tampubolon
    ```


## Cara Menjalankan Aplikasi

1.  **Build Aplikasi:**
    Gunakan perintah `go mod tiny` untuk membangun aplikasi:
    ```bash
    go mod tiny
    ```
.

2.  **Jalankan Aplikasi:**
    Setelah berhasil di-build, jalankan aplikasi dengan perintah:
    ```bash
    go run .\main.go
    ```

3.  **Akses Aplikasi:**
    Setelah aplikasi berjalan, Anda dapat mengakses endpoint API yang tersedia sesuai dengan dokumentasi yang ada di README utama. Biasanya, aplikasi akan berjalan di `http://localhost:8081`.

## Informasi Tambahan

- Pastikan Anda telah mengkonfigurasi database di ./config/database.go dan variabel lingkungan yang diperlukan sebelum menjalankan aplikasi.
- Anda dapat melihat dokumentasi API lengkap di README utama untuk mengetahui endpoint yang tersedia.
- Jika Anda mengalami masalah, periksa log aplikasi untuk informasi lebih lanjut.

## Endpoint

##  Authentication Routes

Endpoint telah terbagi antara endpoint yang membutuh JWT token setelah login dengan label PROTECTED dan label bebas

### 1. Register (POST /register) 

Method: POST * URL: http://localhost:8081/register

Content-Type: application/json * Body (raw/JSON):

Body Raw test:

`{

    "username": "testuser",
    
    "password": "testpassword",
    
    "email": "test@example.com"
    
}`

### 2. Login (POST /login) 

Method: POST * URL: http://localhost:8081/login  

Headers:

Content-Type: application/json * Body (raw/JSON):
Body Raw test:

`{

    "username": "testuser",
    
    "password": "testpassword"
    
}`

## Checklist Routes (Protected)

### 1. Get All Checklists (GET /checklists) 

Method: GET * URL: http://localhost:8081/checklists

Example Response:

[
  {
  
      "id": 1,
      
      "user_id": 1,
      
      "item_name": "My First Checklist",
      
      "created_at": "2024-01-27 12:00:00",
      
      "updated_at": "2024-01-27 12:00:00"
      
  }
  
]

### 2. Create Checklist (POST /checklists)  

Method: POST * URL: http://localhost:8081/checklists
Body Raw test:

`
{

  "name": "My New Checklist"

}
`

### 3. Get Checklist (GET /checklists/{id}) 

Method: GET * URL: http://localhost:8081/checklists/1

 Example Response:
 '
 {
 
    "id": 1,
    
    "user_id": 1,
    
    "item_name": "My First Checklist",
    
    "created_at": "2024-01-27 12:00:00",
    
    "updated_at": "2024-01-27 12:00:00"
    
}
'

### 4. Delete Checklist (DELETE /checklists/{id}

Method: DELETE * URL: http://localhost:8081/checklists/1

## 3. Item Routes (Protected)

### 1. Get Items (GET /checklists/{id}/items) 

Method: GET * URL: http://localhost:8081/checklists/1/items

Example Response:
`
[
  {
      "id": 1,
      "checklist_id": 1,
      "text": "First Item",
      "completed": false,
      "created_at": "2024-01-27 12:00:00",
      "updated_at": "2024-01-27 12:00:00"
  }
]
`

### 2. Create Item (POST /checklists/{id}/items)  

Method: POST * URL: http://localhost:8081/checklists/1/items

Body Raw test:

`
{
    
  "itemName": "My New Item"

}
`



### 3. Update Item (PUT /checklists/{id}/items/{item_id})  

Method: PUT * URL: http://localhost:8081/checklists/1/items/1

{
    "text": "Updated Item Text",
    "completed": true
}

### 4. Delete Item (DELETE /checklists/{id}/items/{item_id}) 

Method: DELETE * URL: http://localhost:8081/checklists/1/items/1


END 