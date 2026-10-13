# Go Project Folder Structure

Struktur folder pada project ini dirancang untuk memisahkan tanggung jawab setiap komponen aplikasi agar kode lebih terorganisir, mudah dipelihara (*maintainable*), dan mudah dikembangkan.

## 📁 Directory Structure

```text
project-name/
├── cmd/
│   └── main.go
├── internal/
│   ├── configs/
│   ├── constants/
│   ├── handlers/
│   ├── model/
│   ├── repository/
│   └── service/
├── scripts/
│   ├── migrations/
├── go.mod
└── go.sum
```

## 📌 Folder Description

### 1. `cmd/`

Folder ini berisi entry point aplikasi. File `main.go` merupakan titik awal eksekusi program yang digunakan untuk menginisialisasi konfigurasi, koneksi database, dependency, dan menjalankan aplikasi.

### 2. `internal/configs/`

Folder ini digunakan untuk menyimpan konfigurasi aplikasi, seperti pengaturan database, environment variables, API keys, dan konfigurasi lainnya.

Pastikan informasi sensitif, seperti password dan secret keys, disimpan melalui environment variables atau mekanisme pengelolaan rahasia yang aman.

### 3. `internal/constants/`

Folder ini digunakan untuk menyimpan konstanta (*constants*) yang digunakan di berbagai bagian aplikasi, seperti kode status, pesan default, nama role, dan nilai tetap lainnya.

### 4. `internal/handlers/`

Sebagai entry point untuk rest apinya, jadi nanti dari main.go itu akan menuju ke handler

### 5. `internal/model/`

Folder ini digunakan untuk mendefinisikan struktur data aplikasi, seperti Go structs, entity, dan model yang berhubungan dengan representasi data serta database.

### 6. `internal/repository/`

Folder ini berisi kode yang bertanggung jawab atas akses dan pengelolaan data dari sumber eksternal, seperti database, Redis, dan penyimpanan lainnya.

Repository menangani operasi data, seperti `Create`, `Find`, `Update`, dan `Delete`, sehingga proses akses data terpisah dari business logic.

### 7. `internal/service/`

Folder ini digunakan untuk menyimpan business logic aplikasi, termasuk validasi bisnis, pemrosesan data, dan koordinasi antar-repository atau komponen lain yang dibutuhkan.

Service layer bertanggung jawab menjalankan aturan bisnis aplikasi agar handler tetap sederhana dan mudah dipelihara.

### 8. `scripts/migrations/`

digunakan untuk menyimpan file-file hasil dari generate Golang Migrate

### . `go.mod`

File `go.mod` berada di root project dan berisi informasi modul Go, termasuk nama module, versi Go yang digunakan, serta dependensi utama yang dibutuhkan oleh project.

### . `go.sum`

File `go.sum` berisi checksum untuk memverifikasi integritas versi modul dan dependensi yang digunakan dalam project Go.

File ini membantu memastikan dependensi yang diunduh sesuai dengan checksum yang tercatat.

## 🔄 Running Application & Flow

go run cmd/main.go

## 🎯 Benefits

> **Note:** Struktur folder ini merupakan salah satu pendekatan yang umum digunakan dalam pengembangan aplikasi Go. Struktur dapat disesuaikan dengan kebutuhan dan kompleksitas project.
