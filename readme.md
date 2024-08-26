# GWT - Go Web Template

GWT (Go Web Template) adalah template awal untuk membangun aplikasi web menggunakan Go. Ini menyediakan struktur dasar untuk server web dengan integrasi database dan manajemen konfigurasi.

## Fitur

- Setup server web dasar
- Integrasi database PostgreSQL
- Manajemen konfigurasi menggunakan variabel lingkungan

## Memulai

### Prasyarat

- Go 1.18 atau lebih baru
- Database PostgreSQL
- File `.env` untuk konfigurasi

### Instalasi

1. **Clone repositori:**

    ```sh
    git clone https://github.com/ekokurniawann/gwt.git
    cd gwt
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Buat file `.env`:**

    Buat file bernama `.env` di direktori root proyek Anda dan tambahkan konfigurasi sesuai kebutuhan. Format file `.env` harus seperti berikut:

    ```env
    DB_HOST=host_database
    DB_PORT=port_database
    DB_USER=user_database
    DB_PASSWORD=password_database
    DB_NAME=nama_database
    SERVER_PORT=:port_server
    ```

4. **Jalankan aplikasi:**

    ```sh
    go run main.go
    ```

    Server akan mulai dan mendengarkan di port yang ditentukan dalam file `.env`.

## Penggunaan

Setelah server berjalan, Anda dapat membuat permintaan HTTP ke endpoint yang didefinisikan dalam aplikasi.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan fork repositori dan kirimkan pull request dengan perubahan Anda.
