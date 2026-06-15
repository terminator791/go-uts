# Go UTS - API Peminjaman

Project ini menyediakan REST API sederhana untuk data peminjaman dengan arsitektur domain-usecase-handler dan koneksi PostgreSQL menggunakan pgx.

## Prasyarat

- Go 1.22+
- PostgreSQL berjalan di `localhost:5432`
- Database: `uts_peminjaman`

## Konfigurasi Environment

Aplikasi menggunakan environment variable berikut (default ditunjukkan):

- `DB_HOST=localhost`
- `DB_PORT=5432`
- `DB_USER=postgres`
- `DB_PASSWORD=postgres`
- `DB_NAME=uts_peminjaman`
- `SERVER_ADDR=:8080`

## Perintah Penting

Gunakan Makefile untuk menjalankan perintah:

- `make migrate` : menjalankan migrasi schema
- `make fresh` : drop semua tabel lalu migrate ulang
- `make seed` : isi data awal
- `make run` : menjalankan server

## Endpoint

- `POST /api/anggota`
- `GET /api/anggota`
- `GET /api/anggota/{id}`
- `PUT /api/anggota/{id}`
- `DELETE /api/anggota/{id}`
- `POST /api/peminjaman`
- `GET /api/peminjaman`
- `GET /api/peminjaman/{id}`
- `PUT /api/peminjaman/{id}`
- `DELETE /api/peminjaman/{id}`
- `POST /api/pembayaran`
- `GET /api/pembayaran`
- `GET /api/pembayaran/{id}`
- `PUT /api/pembayaran/{id}`
- `GET /api/peminjaman/{id}/pembayaran`
- `DELETE /api/pembayaran/{id}`

Detail request/response lihat di `REQUIREMENTS.md`.

## Curl Testing

Contoh curl untuk seluruh endpoint tersedia di `CURL.md`.
