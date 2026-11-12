## 5. Kontrol akses dan isolasi rumah tangga

Aturan siapa melihat apa kanonis di 06 §9 dan 07 §5.3. Bagian ini menetapkan cara menegakkannya sehingga dapat diuji.

| ID | Persyaratan | Untuk | Tahap | Label | Cara menguji |
|---|---|---|---|---|---|
| SEC-01 | Setiap perintah dan setiap bacaan diotorisasi di server dari keanggotaan, peran, dan himpunan penerima saat itu, bukan dari yang dikirim klien. Klien tidak pernah menerima data yang tidak boleh ia tampilkan untuk "disembunyikan di layar" | 09 | Pilot | `ENGINEERING-STANDARDS.md` §4; **[Usulan]** | Untuk setiap jenis objek dan tingkat visibilitas: panggil API sebagai setiap peran; tanggapan tidak memuat medan yang tidak boleh dibaca, juga sebagai `null` yang membedakan "ada tapi tertutup" dari "tidak ada" |
| SEC-02 | Setiap baris membawa `household_id`; setiap indeks unik dan foreign key atas data rumah tangga diawali `household_id`, sehingga pelanggaran constraint tidak membuka keberadaan baris rumah tangga lain (18 §5.5 butir 4). Pesan galat constraint tidak diteruskan ke klien | 09 | Pilot | **[Usulan]** | Uji FR-FAM-020/AC-1 diperluas: kirim pengenal, kunci idempotency, dan nilai unik yang sudah dipakai rumah tangga lain; tanggapan identik dengan nilai yang belum pernah dipakai |
