1. [A000124 – Sloane's OEIS Sequence](#1-a000124--sloanes-oeis)
2. [Dense Ranking](#2-dense-ranking)
3. [Balanced Bracket](#3-balanced-bracket)

---

## 1. A000124 – Sloane’s OEIS

**Deskripsi**  
Menghasilkan deret berdasarkan rumus:

```
a(n) = n(n+1)/2 + 1
```

**Contoh Output**:
- Input: `5` → Output: `[1 2 4 7 11]`
- Input: `7` → Output: `[1 2 4 7 11 16 22]`
- Input: `10` → Output: `[1 2 4 7 11 16 22 29 37 46]`

**Kompleksitas:**
- **Waktu**: O(n)
- **Ruang**: O(n)

---

## 2. Dense Ranking

**Deskripsi**  
Menghitung peringkat GITS dalam leaderboard menggunakan sistem **Dense Ranking**, di mana:
- Skor tertinggi mendapat peringkat 1
- Skor yang sama → peringkat sama
- Skor lebih kecil → peringkat bertambah

**Contoh**:
```
Leaderboard: [100, 100, 50, 40, 40, 20, 10]
GITS Scores: [5, 25, 50, 120]
Output:      [6, 4, 2, 1]
```

**Kompleksitas:**
- **Waktu**: O(m log n), dengan binary search
- **Ruang**: O(n) untuk leaderboard unik

---

## 3. Balanced Bracket

**Deskripsi**  
Memverifikasi keseimbangan string bracket dengan karakter:  
`()`, `{}`, `[]`.

Menggunakan struktur **stack** untuk mencocokkan setiap bracket buka dan tutup.

**Contoh**:
```
Input: {[(])}           → Output: NO
Input: {([{}])[]}       → Output: YES
Input: ((({[()]})[]))   → Output: YES
```

**Kompleksitas:**
- **Waktu**: O(n)
- **Ruang**: O(n)

---

## Cara Menjalankan

1. Pastikan Go telah terinstal.
2. Jalankan program:

```bash
go run main.go
```

3. Program akan menampilkan hasil dari ketiga soal dengan 3 input berbeda untuk masing-masing.

---

## Struktur Program

```plaintext
main.go        # Implementasi semua soal dalam satu file
output.png     # Output dari setiap soal
README.md      # Dokumentasi ini
```

---

