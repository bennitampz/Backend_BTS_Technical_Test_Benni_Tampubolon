﻿# Backend_BTS_Technical_Test_Benni_Tampubolon

## Deskripsi

Repositori ini berisi implementasi backend untuk aplikasi Notes App. Aplikasi ini menyediakan fitur untuk registrasi dan login pengguna, serta pengelolaan checklist dan item-item di dalamnya.

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

# Cara Menjalankan Aplikasi Backend Technical_Test

Dokumen ini menjelaskan langkah-langkah untuk menjalankan aplikasi backend NoteSApp yang dibuat dengan Go.

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

- Pastikan Anda telah mengkonfigurasi database di .env dan variabel lingkungan yang diperlukan sebelum menjalankan aplikasi.
- Anda dapat melihat dokumentasi API lengkap di README utama untuk mengetahui endpoint yang tersedia.
- Jika Anda mengalami masalah, periksa log aplikasi untuk informasi lebih lanjut.

## Endpoint

##  Authentication Routes

Endpoint telah terbagi antara endpoint yang membutuhkan JWT token setelah login dengan label PROTECTED dan label bebas

### 1. Register (POST /register) 

Method: POST * URL: http://localhost:8081/register

Content-Type: application/json * Body (raw/JSON):

Body Raw test:

{

    "username": "testuser",
    
    "password": "testpassword",
    
    "email": "test@example.com"
    
}

### 2. Login (POST /login) 

Method: POST * URL: http://localhost:8081/login  

Headers:

Content-Type: application/json * Body (raw/JSON):
Body Raw test:

{

    "username": "testuser",
    
    "password": "testpassword"
    
}


## Checklist Routes (Protected)

### 1. Get All Checklists (GET /checklists) 

Method: GET + Bearer Token * URL: http://localhost:8081/checklists

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

Method: POST + Bearer Token * URL: http://localhost:8081/checklists
Body Raw test:


{

  "name": "My New Checklist"

}


### 3. Get Checklist (GET /checklists/{id}) 

Method: GET + Bearer Token * URL: http://localhost:8081/checklists/1

 Example Response:
 
 {
 
    "id": 1,
    
    "user_id": 1,
    
    "item_name": "My First Checklist",
    
    "created_at": "2024-01-27 12:00:00",
    
    "updated_at": "2024-01-27 12:00:00"
    
}


### 4. Delete Checklist (DELETE /checklists/{id}

Method: DELETE + Bearer Token * URL: http://localhost:8081/checklists/1

## 3. Item Routes (Protected)

### 1. Get Items (GET /checklists/{id}/items) 

Method: GET + Bearer Token * URL: http://localhost:8081/checklists/1/items

Example Response:

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


### 2. Create Item (POST /checklists/{id}/items)  

Method: POST + Bearer Token * URL: http://localhost:8081/checklists/1/items

Body Raw test:


{
    
  "itemName": "My New Item"

}




### 3. Update Item (PUT /checklists/{id}/items/{item_id})  

Method: PUT + Bearer Token * URL: http://localhost:8081/checklists/1/items/1

{

    "text": "Updated Item Text",
    
    "completed": true
    
}

### 4. Delete Item (DELETE /checklists/{id}/items/{item_id}) 

Method: DELETE + Bearer Token * URL: http://localhost:8081/checklists/1/items/1

1. API login ✓

![image](https://github.com/user-attachments/assets/3ad49c1d-5f1b-4f69-9cad-fd7c85da79f7)
   
2. API daftar baru ✓

![image](https://github.com/user-attachments/assets/9382bca3-f482-4db3-9352-8763598a2342)

5. API untuk membuat checklist (berdasarkan contoh gambar 2, adalah kotak-kotak yang berwarna) ✓

![image](https://github.com/user-attachments/assets/baa52efe-3cac-451a-9fe8-2b03b0693370)

6. API untuk menghapus checklist ✓

![image](https://github.com/user-attachments/assets/035f1264-6e0a-43b2-a1b2-9a25f68a3231)

8. API untuk menampilkan checklist-checklist yang sudah dibuat ✓

![image](https://github.com/user-attachments/assets/b6e8fa88-f74f-4a47-9228-635549aa0ba0)

10. API Detail Checklist (Berisi item-item to-do yang sudah dibuat) ✓

![image](https://github.com/user-attachments/assets/33a258ae-0f63-4acd-ba8e-931903b29c6c)
    
12. API untuk membuat item-item to-do di dalam checklist ✓

![image](https://github.com/user-attachments/assets/de3b50a1-6a3e-407b-8fcb-33c1914079dc)

14. API detail item ✓

![image](https://github.com/user-attachments/assets/a8bc5bf9-daf6-45b8-b9e5-f7ddb416e5f5)


16. API untuk mengubah item-item di dalam checklist ✓

![image](https://github.com/user-attachments/assets/4ac47768-5774-4d5b-b0ae-93187181857f)

18. API untuk mengubah status dari item di dalam checklist (misalkan item sudah selesai dilakukan) ✓

![image](https://github.com/user-attachments/assets/7865e01d-2801-4a18-b250-ab248a861156)

20. API untuk menghapus item dari checklist ✓

![image](https://github.com/user-attachments/assets/8e589b3d-36c9-4a45-8090-5784e5fd4271)

END 
