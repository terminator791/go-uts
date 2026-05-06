CREATE TABLE IF NOT EXISTS anggota (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    alamat VARCHAR(255) NOT NULL,
    telepon VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS peminjaman (
    id SERIAL PRIMARY KEY,
    id_anggota INT NOT NULL REFERENCES anggota(id) ON DELETE CASCADE,
    tanggal_pinjam DATE NOT NULL,
    jumlah_pinjam BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS pembayaran (
    id SERIAL PRIMARY KEY,
    id_peminjaman INT NOT NULL REFERENCES peminjaman(id) ON DELETE CASCADE,
    tanggal_pembayaran DATE NOT NULL,
    jumlah_pembayaran BIGINT NOT NULL
);
