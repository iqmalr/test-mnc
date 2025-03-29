# **Project Installment & Payment API**

Proyek ini adalah API sederhana untuk kebutuhan test MNC menggunakan **Golang (Gin)** untuk mengelola transaksi pinjaman uang (_installment_), pembayaran cicilan (_payment_), dan rekap cicilan (_recap_).

## **Demo Project**

Saya telah memindahkan kode ke **Project IDX**, sehingga Anda dapat mencoba demo API tanpa perlu meng-clone proyek ke lokal.

Silakan kunjungi link berikut untuk mengakses **Swagger UI** dan mencoba API secara langsung:

🔗 **[Demo API - Swagger UI](https://8082-idx-test-mnc-1743261436111.cluster-7ubberrabzh4qqy2g4z7wgxuw2.cloudworkstations.dev/swagger/index.html#/)**

## **Fitur Utama**

- **Create Installment** → Membuat transaksi pinjaman uang.
- **Payment** → Melakukan cicilan terhadap pinjaman yang dibuat.
- **Rekap Cicilan** → Melihat status dan histori pembayaran cicilan.
- **User & Merchant Management** → Melihat daftar pengguna dan merchant.
- **Authentication** → Menggunakan middleware untuk proteksi API.

---

## **Endpoint API**

### **1. Authentication**

| Method   | Endpoint | Description                                |
| -------- | -------- | ------------------------------------------ |
| **POST** | `/login` | Login untuk mendapatkan token autentikasi. |

---

### **2. Installment (Cicilan/Peminjaman)**

| Method   | Endpoint       | Description                      |
| -------- | -------------- | -------------------------------- |
| **POST** | `/installment` | Membuat transaksi pinjaman baru. |
| **GET**  | `/installment` | Menampilkan semua data pinjaman. |

---

### **3. Payment (Pembayaran Cicilan)**

| Method   | Endpoint    | Description                                        |
| -------- | ----------- | -------------------------------------------------- |
| **POST** | `/payments` | Melakukan pembayaran cicilan.                      |
| **GET**  | `/payments` | Menampilkan semua pembayaran yang telah dilakukan. |

---

### **4. Recap (Rekap Cicilan)**

| Method  | Endpoint | Description                                       |
| ------- | -------- | ------------------------------------------------- |
| **GET** | `/recap` | Melihat status cicilan dan histori pembayarannya. |

---

### **5. User & Merchant**

| Method  | Endpoint     | Description                        |
| ------- | ------------ | ---------------------------------- |
| **GET** | `/users`     | Menampilkan daftar semua pengguna. |
| **GET** | `/merchants` | Menampilkan daftar semua merchant. |

---

## **Cara Menjalankan Project**

1. **Clone repository dari GitHub:**

   ```sh
   git clone https://github.com/iqmalr/test-mnc
   cd test-mnc
   ```

2. **Install dependencies:**

   ```sh
   go mod tidy
   ```

3. **Jalankan project Go:**

   - Menggunakan **Air** (live reload):
     ```sh
     air
     ```
   - Atau menjalankan secara manual:
     ```sh
     go run main.go
     ```

4. **API bisa diakses di:**
   ```sh
   http://localhost:8082
   ```

---

## **Mengakses API Documentation**

1. **Swagger UI tersedia di:**
   ```
   http://localhost:8082/swagger/index.html#
   ```
2. **Login terlebih dahulu:**
   - Anda bisa menggunakan akun yang tersedia di **`data/user.json`**
   - Contoh akun yang bisa digunakan:
     ```json
     {
       "email": "iqmalr@gmail.com",
       "password": "password"
     }
     ```

---

## **Menggunakan Fitur Pinjaman**

### **1. Membuat Pinjaman Baru**

1. **Navigasi ke endpoint** `POST /installment` **di Swagger UI**
2. **Masukkan data dalam format berikut:**

   ```json
   {
     "user_id": 1,
     "merchant_id": 2,
     "total_amount": 5000000
   }
   ```

   - `user_id`: ID user yang mengajukan pinjaman
   - `merchant_id`: ID merchant tempat pinjaman dibuat
   - `total_amount`: Jumlah total pinjaman

---

### **2. Melakukan Pembayaran Cicilan**

1. **Navigasi ke endpoint** `POST /payments` **di Swagger UI**
2. **Masukkan data pembayaran dalam format berikut:**

   ```json
   {
     "transaction_id": 1,
     "amount": 200000,
     "payment_method": "bank_transfer"
   }
   ```

   - `transaction_id`: ID transaksi pinjaman yang ingin dibayar
   - `amount`: Jumlah pembayaran (dalam Rupiah)
   - `payment_method`: Metode pembayaran (contoh: `"bank_transfer"`)

---

## **Catatan Tambahan**

- **Data pengguna, merchant, installment, dan payment disimpan dalam file JSON** di dalam folder **`data/`** (digunakan sebagai penyimpanan sementara).
- **Semua endpoint yang memerlukan autentikasi harus diakses setelah melakukan login.**
- **Pastikan server berjalan sebelum mengakses Swagger UI.**

---

🔥 **Proyek ini menggunakan `Gin` sebagai web framework dan `JWT Middleware` untuk proteksi API.**
