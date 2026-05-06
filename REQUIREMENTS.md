# Dokumentasi Requirement API Peminjaman

Dokumen ini berisi spesifikasi kebutuhan untuk pembuatan RESTful API Sistem Peminjaman beserta struktur database yang dibutuhkan berdasarkan soal UTS.

## 1. Struktur Database

Terdapat 3 tabel yang perlu dibuat dalam database:

### Tabel `anggota`

| Kolom     | Tipe Data (Estimasi) | Keterangan                        |
| --------- | -------------------- | --------------------------------- |
| `id`      | INT (Primary Key)    | ID unik anggota                   |
| `nama`    | VARCHAR              | Nama anggota                      |
| `alamat`  | TEXT / VARCHAR       | Alamat tempat tinggal             |
| `telepon` | VARCHAR              | Nomor telepon yang bisa dihubungi |

_Contoh Data:_

- `1` | `John Doe` | `Jalan Merdeka No. 123` | `081234567890`
- `2` | `Jane Smith` | `Jalan Gatot Subroto` | `082345678901`

### Tabel `peminjaman`

| Kolom            | Tipe Data (Estimasi) | Keterangan                                       |
| ---------------- | -------------------- | ------------------------------------------------ |
| `id`             | INT (Primary Key)    | ID unik peminjaman                               |
| `id_anggota`     | INT (Foreign Key)    | Merujuk ke `id` tabel `anggota`                  |
| `tanggal_pinjam` | DATE                 | Tanggal saat peminjaman dilakukan                |
| `jumlah_pinjam`  | DECIMAL/INT          | Nominal uang yang dipinjam                       |
| `status`         | VARCHAR/ENUM         | Status peminjaman (contoh: `pending`, `selesai`) |

_Contoh Data:_

- `1` | `1` | `2024-04-10` | `1000000` | `pending`
- `2` | `2` | `2024-04-12` | `750000` | `selesai`

### Tabel `pembayaran`

| Kolom                | Tipe Data (Estimasi) | Keterangan                         |
| -------------------- | -------------------- | ---------------------------------- |
| `id`                 | INT (Primary Key)    | ID unik pembayaran                 |
| `id_peminjaman`      | INT (Foreign Key)    | Merujuk ke `id` tabel `peminjaman` |
| `tanggal_pembayaran` | DATE                 | Tanggal pembayaran dilakukan       |
| `jumlah_pembayaran`  | DECIMAL/INT          | Nominal uang yang dibayarkan       |

_Contoh Data:_

- `1` | `1` | `2024-04-15` | `500000`
- `2` | `1` | `2024-04-20` | `500000`

---

## 2. API Endpoints

Sistem wajib mengimplementasikan fungsionalitas CRUD form API JSON untuk entitas `peminjaman`.

### 1. Membuat entri baru (Create)

Tugas: Buatlah sebuah API endpoint untuk membuat (Create) entri baru dalam database. Pastikan API endpoint menerima input dalam format JSON dan menyimpan data yang diterima ke dalam database.

- **Method:** `POST`
- **Endpoint:** `/api/peminjaman`
- **Request Body:** (JSON)
  ```json
  {
    "id_anggota": 1,
    "tanggal_pinjam": "2024-04-16",
    "jumlah_pinjam": 500000,
    "status": "pending"
  }
  ```
- **Response:** (JSON)
  ```json
  {
    "message": "Data peminjaman berhasil ditambahkan",
    "data": {
      "id": 1,
      "id_anggota": 1,
      "tanggal_pinjam": "2024-04-16",
      "jumlah_pinjam": 500000,
      "status": "pending"
    }
  }
  ```

### 2. Mengambil semua entri (Read)

Tugas: Implementasikan fungsi untuk mengambil (Read) semua entri dari database dan mengembalikannya dalam format JSON.

- **Method:** `GET`
- **Endpoint:** `/api/peminjaman`
- **Response:** (JSON)
  ```json
  {
    "data": [
      {
        "id": 1,
        "id_anggota": 1,
        "tanggal_pinjam": "2024-04-16",
        "jumlah_pinjam": 500000,
        "status": "pending"
      },
      {
        "id": 2,
        "id_anggota": 2,
        "tanggal_pinjam": "2024-04-15",
        "jumlah_pinjam": 750000,
        "status": "selesai"
      }
    ]
  }
  ```

### 3. Mengambil entri berdasarkan ID (Read)

Tugas: Buat API endpoint untuk mengambil (Read) sebuah entri berdasarkan ID yang diberikan dalam format JSON.

- **Method:** `GET`
- **Endpoint:** `/api/peminjaman/{id}`
- **Response:** (JSON)
  ```json
  {
    "data": {
      "id": 1,
      "id_anggota": 1,
      "tanggal_pinjam": "2024-04-16",
      "jumlah_pinjam": 500000,
      "status": "pending"
    }
  }
  ```

### 4. Mengupdate entri berdasarkan ID (Update)

Tugas: Implementasikan fungsi untuk mengupdate (Update) entri dalam database berdasarkan ID. Menerima format JSON dan memperbarui data di database yang sesuai.

- **Method:** `PUT`
- **Endpoint:** `/api/peminjaman/{id}`
- **Request Body:** (JSON)
  ```json
  {
    "status": "selesai"
  }
  ```
- **Response:** (JSON)
  ```json
  {
    "message": "Data peminjaman berhasil diperbarui",
    "data": {
      "id": 1,
      "id_anggota": 1,
      "tanggal_pinjam": "2024-04-16",
      "jumlah_pinjam": 500000,
      "status": "selesai"
    }
  }
  ```

### 5. Menghapus entri berdasarkan ID (Delete)

Tugas: Buat API endpoint untuk menghapus (Delete) entri dari database berdasarkan ID yang diberikan.

- **Method:** `DELETE`
- **Endpoint:** `/api/peminjaman/{id}`
- **Response:** (JSON)
  ```json
  {
    "message": "Data peminjaman berhasil dihapus"
  }
  ```
