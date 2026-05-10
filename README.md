# Ice Sliding Puzzle Solver (Web UI)

## a. Penjelasan Singkat Program
Ice Sliding Puzzle Solver adalah program berbasis **Web GUI** yang dibangun menggunakan bahasa Go (Golang) sebagai *backend* dan HTML/JS/Tailwind CSS sebagai *frontend*. Program ini mengimplementasikan algoritma *Pathfinding* (UCS, GBFS, A*, dan BFS sebagai bonus) untuk memecahkan permainan *Ice Sliding Puzzle*. 

Dalam permainan ini, sebuah pin/aktor (`Z`) harus meluncur di atas es licin secara vertikal atau horizontal sampai berhenti saat menabrak rintangan (`X`). Aktor juga harus melewati serangkaian angka secara berurutan sebelum mencapai titik keluar/tujuan (`O`), tanpa melewati lava (`L`). Program ini menampilkan rute solusi, total *cost*, waktu eksekusi, serta jumlah iterasi yang dilakukan, dan dilengkapi dengan fitur *Interactive Playback* berbentuk antarmuka web yang mulus untuk melihat proses penyelesaian tahap demi tahap.

## b. Struktur Direktori
```text
.
├── src/
│   ├── go.mod
│   ├── go.sum
│   ├── main.go           <-- Entry point Web Server (Backend)
│   ├── static/           <-- Folder Frontend
│   │   ├── index.html    <-- Layout Web UI
│   │   └── app.js        <-- Logika API & Animasi Playback
│   ├── model/            <-- Logika Solver & Algoritma
│   │   ├── direction.go
│   │   ├── state.go
│   │   ├── priorityQueue.go
│   │   ├── solver.go
│   │   ├── algorithms.go
│   │   └── heuristic.go
│   └── utils/
│       └── parser.go     <-- Parser map dari upload memori
├── test/                 <-- Folder contoh file input (.txt)
├── docs/  
└── README.md
```

## c. Requirement Program dan Instalasi
Untuk dapat menjalankan program ini, Anda hanya memerlukan:
1. **Go (Golang)** terpasang pada sistem Anda. Anda dapat mengunduhnya di [golang.org](https://go.dev/).
2. Browser modern (Chrome, Edge, Firefox, Safari).

*Catatan: Aplikasi ini dirancang agar bersih. Semua dibangun murni menggunakan fungsi bawaan Go (Standard Library) tanpa ada dependency pihak ketiga yang memberatkan.*

## d. Cara Mengkompilasi Program
Program ini dapat dikompilasi menjadi *executable binary file* server web tunggal.
1. Buka terminal atau command prompt.
2. Pindah ke direktori kode sumber (`src`):
   ```bash
   cd src
   ```
3. Lakukan kompilasi menggunakan perintah `go build`:
   ```bash
   go build -o puzzle_solver.exe .
   ```

## e. Cara Menjalankan dan Menggunakan Program
Sangat disarankan untuk menjalankan program ini langsung melalui `go run` saat proses pengujian.

**Cara Menjalankan Server:**
1. Buka terminal, pastikan Anda berada di folder `src`.
   ```bash
   cd src
   ```
2. Jalankan perintah:
   ```bash
   go run main.go
   ```
3. Buka web browser Anda dan ketikkan alamat: **`http://localhost:8080`**

**Langkah-Langkah Penggunaan:**
1. Di halaman web, unggah file papan permainan (`.txt`) dengan cara mengklik kolom input file atau melakukan *Drag & Drop*.
2. Pilih algoritma pencarian rute yang diinginkan (UCS, GBFS, A*, atau BFS).
3. Jika menggunakan algoritma `GBFS` atau `A*`, pilih salah satu opsi heuristik:
   - `H1`: Manhattan Distance
   - `H2`: Euclidean Distance
   - `H3`: Chebyshev Distance
   - `H4`: Missing Targets (Jumlah target tersisa)
   - `H5`: Straight to Goal (Jarak lurus ke target)
4. Tekan tombol **Solve Puzzle**.
5. Setelah rute ditemukan, *panel* hasil di sebelah kiri akan memunculkan metrik waktu eksekusi dan iterasi.
6. Anda bisa memutar visualisasi rute dengan menekan tombol Play (▶️) di bawah papan permainan, memutarnya secara manual dengan (⏭️/⏮️), atau melompat ke step tertentu menggunakan *Slider*.
7. Klik tombol **Download Log Solusi (.txt)** untuk mengunduh log riwayat lengkap langsung ke komputer Anda.

## f. Author
- **Wisa Ahmaduta Dinutama / 18223003**
- **Muhammad Adam Mirza / 18223015**
