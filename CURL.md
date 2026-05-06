# Curl Test Commands

Gunakan perintah berikut untuk menguji seluruh endpoint. Pastikan server berjalan di `http://localhost:8080`.

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

## 9. Read Pembayaran By Peminjaman (GET /api/peminjaman/{id}/pembayaran)

```sh
curl http://localhost:8080/api/peminjaman/1/pembayaran
```

## 10. Delete Pembayaran (DELETE /api/pembayaran/{id})

```sh
curl -X DELETE http://localhost:8080/api/pembayaran/1
```
