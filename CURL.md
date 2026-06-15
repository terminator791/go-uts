# Curl Test Commands

Gunakan perintah berikut untuk menguji seluruh endpoint. Pastikan server berjalan di `http://localhost:8080`.

## 0. Create Anggota (POST /api/anggota)

```sh
curl -X POST http://localhost:8080/api/anggota \
  -H 'Content-Type: application/json' \
  -d '{"nama":"John Doe","alamat":"Jalan Merdeka No. 123","telepon":"081234567890"}'
```

## 0.1 Read All Anggota (GET /api/anggota)

```sh
curl http://localhost:8080/api/anggota
```

## 0.2 Read Anggota By ID (GET /api/anggota/{id})

```sh
curl http://localhost:8080/api/anggota/1
```

## 0.3 Update Anggota (PUT /api/anggota/{id})

```sh
curl -X PUT http://localhost:8080/api/anggota/1 \
  -H 'Content-Type: application/json' \
  -d '{"nama":"Jane Smith","alamat":"Jalan Gatot Subroto","telepon":"082345678901"}'
```

## 0.4 Delete Anggota (DELETE /api/anggota/{id})

```sh
curl -X DELETE http://localhost:8080/api/anggota/1
```

## 1. Create (POST /api/peminjaman)

```sh
curl -X POST http://localhost:8080/api/peminjaman \
  -H 'Content-Type: application/json' \
  -d '{"id_anggota":1,"tanggal_pinjam":"2024-04-16","jumlah_pinjam":500000,"status":"pending"}'
```

## 2. Read All (GET /api/peminjaman)

```sh
curl http://localhost:8080/api/peminjaman
```

## 3. Read By ID (GET /api/peminjaman/{id})

```sh
curl http://localhost:8080/api/peminjaman/1
```

## 4. Update (PUT /api/peminjaman/{id})

```sh
curl -X PUT http://localhost:8080/api/peminjaman/1 \
  -H 'Content-Type: application/json' \
  -d '{"status":"selesai"}'
```

## 5. Delete (DELETE /api/peminjaman/{id})

```sh
curl -X DELETE http://localhost:8080/api/peminjaman/1
```

## 6. Create Pembayaran (POST /api/pembayaran)

```sh
curl -X POST http://localhost:8080/api/pembayaran \
  -H 'Content-Type: application/json' \
  -d '{"id_peminjaman":1,"tanggal_pembayaran":"2024-04-15","jumlah_pembayaran":500000}'
```

## 7. Read All Pembayaran (GET /api/pembayaran)

```sh
curl http://localhost:8080/api/pembayaran
```

## 8. Read Pembayaran By ID (GET /api/pembayaran/{id})

```sh
curl http://localhost:8080/api/pembayaran/1
```

## 9. Update Pembayaran (PUT /api/pembayaran/{id})

```sh
curl -X PUT http://localhost:8080/api/pembayaran/1 \
  -H 'Content-Type: application/json' \
  -d '{"id_peminjaman":1,"tanggal_pembayaran":"2024-04-16","jumlah_pembayaran":250000}'
```

## 10. Read Pembayaran By Peminjaman (GET /api/peminjaman/{id}/pembayaran)

```sh
curl http://localhost:8080/api/peminjaman/1/pembayaran
```

## 11. Delete Pembayaran (DELETE /api/pembayaran/{id})

```sh
curl -X DELETE http://localhost:8080/api/pembayaran/1
```
