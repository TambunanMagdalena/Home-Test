-- Buat tabel fields sesuai kebutuhan aplikasimu. Tambahkan kolom lain yang diperlukan oleh aplikasi.
CREATE TABLE IF NOT EXISTS fields (
  id SERIAL PRIMARY KEY,
  name TEXT NOT NULL,
  description TEXT,
  location TEXT,
  price numeric(12,2),
  capacity integer DEFAULT 0,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
);