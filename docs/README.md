# Family OS — Dokumentasi Perencanaan

Pintu masuk seluruh dokumen perencanaan Family OS. Fase saat ini adalah **perencanaan saja**: belum ada kode aplikasi.

## Tujuan

Suami dan istri mengetahui keadaan uang rumah tangga dan urusan rumah tangga tanpa harus saling bertanya, dengan usaha pencatatan serendah mungkin, dan aplikasi membantu mereka saling memperhatikan dan bersikap hangat satu sama lain.

Tujuan ini bertahap: pertama rumah tangga pilot (rumah tangga pemilik produk), kemudian produk publik untuk keluarga Indonesia. Setiap dokumen dan keputusan harus dapat ditelusuri ke tujuan ini dan ke masalah inti di `family-os-brief.md` §1.1.

## Sumber kebenaran

Urutan dari yang tertinggi. Bila dua sumber bertentangan, ikuti yang lebih tinggi dan catat pertentangannya di [17-open-questions-assumptions.md](17-open-questions-assumptions.md).

1. [`family-os-brief.md`](../family-os-brief.md) — arah, cakupan, tingkatan MVP, tahapan, checkpoint.
2. [00-glossary.md](00-glossary.md) — satu-satunya definisi setiap istilah bersama.
3. [`adr/`](adr/) — keputusan yang sudah dibuat.
4. [`ENGINEERING-STANDARDS.md`](../ENGINEERING-STANDARDS.md) — aturan keras fase pembangunan.
5. Dokumen lain di `docs/`.

## Urutan baca

Pembaca baru: mulai dari [00-glossary.md](00-glossary.md), lalu [_progress.md](_progress.md) untuk mengetahui posisi terakhir, lalu dokumen sesuai tahap. Ringkasan eksekutif (01) ditulis paling akhir dan menjadi titik masuk setelah tahap S5 selesai.

## Daftar dokumen

Status: **belum ada** → **draf** → **ditinjau** → **final**. Status terkini selalu dicatat di [_progress.md](_progress.md); tabel ini hanya menunjukkan tahap asal tiap dokumen.

| # | Dokumen | Isi | Tahap | Pemilik |
|---|---|---|---|---|
| — | [_progress.md](_progress.md) | Log tahapan, keputusan, asumsi, pertanyaan terbuka, tahap berikutnya | Semua | lead |
| 00 | [Glosarium](00-glossary.md) | Definisi kanonis setiap istilah bersama | S1 (istilah keuangan difinalkan di S2) | lead; bagian istilah keuangan oleh `finance-domain` |
| 01 | [Ringkasan eksekutif](01-executive-summary.md) | Visi, masalah inti, audiens, proposisi nilai, diferensiasi, MVP dalam satu halaman | S5 (terakhir) | lead |
| 02 | [Prinsip produk](02-product-principles.md) | Prinsip desain, UX, privasi, dan integritas keuangan yang dapat diuji | S2 | `product-analyst` |
| 03 | [Persona dan model rumah tangga](03-personas-household-models.md) | Persona, struktur rumah tangga, kebutuhan inti | S2 | `product-analyst` |
| 04 | [Peta fitur](04-feature-map.md) | Modul, fitur, sub-fitur, hubungan, dependensi | S3 | `product-analyst` |
| 05 | [Kebutuhan fungsional](05-functional-requirements.md) | Spesifikasi, business rule, izin, alur, kriteria penerimaan | S3 | `product-analyst` |
| 06 | [Arsitektur aktivitas dan persetujuan](06-activity-approval-architecture.md) | Model alur kerja, state machine, logika persetujuan, settlement | S2 | `architect` |
| 07 | [Domain keuangan](07-financial-domain.md) | Konsep ledger, saldo, custody, anggaran, transaksi, laporan, contoh terhitung | S2 | `finance-domain` |
| 08 | [Strategi UX/UI](08-ux-ui-strategy.md) | Arsitektur informasi, navigasi, layar, perjalanan pengguna, design system | S4 | `ux-designer` |
| 09 | [Arsitektur teknis](09-technical-architecture.md) | Perbandingan, stack yang direkomendasikan, komponen, data, realtime, offline | S4 | `architect` |
| 10 | [Keamanan dan privasi](10-security-privacy.md) | Threat model, kontrol akses, data sensitif, pencegahan penyalahgunaan, siklus hidup data, regulasi | S4 | `security-privacy` |
| 11 | [Inovasi](11-innovation.md) | ≥ 15 konsep fitur beserta penilaian dan prioritas | S3 | `product-analyst` |
| 12 | [Skenario](12-scenarios.md) | ≥ 20 skenario end-to-end beserta kasus tepi | S3 | `product-analyst` |
| 13 | [Roadmap produk](13-product-roadmap.md) | Fase, rencana validasi, milestone, dependensi, model bisnis | S5 | `product-analyst` |
| 14 | [Roadmap engineering](14-engineering-roadmap.md) | Workstream, strategi pengujian, quality gate, rilis, operasi | S5 | `architect` |
| 15 | [Penilaian risiko](15-risk-assessment.md) | Risiko produk, UX, keamanan, keuangan, teknis, bisnis beserta mitigasi | S5 | `architect` |
| 16 | [Rekomendasi akhir](16-final-recommendations.md) | MVP, arsitektur, keputusan kritis, keputusan yang ditunda | S5 | lead |
| 17 | [Pertanyaan terbuka dan asumsi](17-open-questions-assumptions.md) | Keputusan yang menunggu pemilik, asumsi, default yang diusulkan | S1, diperbarui tiap tahap | `product-analyst` |
| 18 | [Riset pasar](18-market-research.md) | Pasar, regulasi, integrasi, desain platform, teknologi — beserta sumber dan tanggal akses | S1, diperluas tiap ada riset baru | `researcher` |
| — | [`adr/`](adr/) | Architecture Decision Record (`ADR-NNN-judul.md`) | Mulai S2 | `architect` |

Tautan ke dokumen yang berstatus **belum ada** baru berfungsi setelah tahap asalnya selesai.

## Tahapan dan checkpoint

| Tahap | Keluaran | Checkpoint |
|---|---|---|
| S1 — Framing | README, 00, 17, 18 | **STOP**: 10 pertanyaan paling menentukan beserta default yang diusulkan |
| S2 — Fondasi domain | 02, 03, 07, 06 | **STOP**: ringkasan model keuangan dan model persetujuan ≤ 1 halaman |
| S3 — Definisi produk | 04, 05, 11, 12 | **STOP**: ringkasan potongan fitur MVP |
| S4 — Pengalaman dan arsitektur | 08, 09 (+ ADR), 10 | **STOP**: stack dan arsitektur yang direkomendasikan beserta tradeoff utama |
| S5 — Rencana dan sintesis | 13, 14, 15, 16, lalu 01 | — |
| S6 — Pemeriksaan konsistensi | Perbaikan kontradiksi, tautan rusak, istilah tak terdefinisi, kebutuhan tak tertelusur | Laporan perubahan |

Tahap berikutnya tidak dimulai tanpa persetujuan eksplisit pemilik produk di setiap **STOP**.

## Konvensi

### Bahasa

Semua dokumen di `docs/` ditulis dalam bahasa Indonesia yang profesional dan wajar. Istilah teknis yang sudah mapan (ledger, offline-first, state machine, RBAC, dan sejenisnya) tetap dalam bahasa Inggris.

### Label bukti

Setiap klaim yang tidak sepele diberi satu label:

| Label | Arti |
|---|---|
| **[Terverifikasi]** | Dikonfirmasi lewat riset; menyertakan sumber dan tanggal akses |
| **[Asumsi]** | Keyakinan yang masuk akal, belum diverifikasi |
| **[Usulan]** | Usulan atau rekomendasi tim perencana |
| **[Perlu validasi]** | Harus dicek ke pengguna, ahli hukum, atau mitra sebelum diandalkan; menyebut apa yang sudah dicari dan mengapa belum bisa dikonfirmasi |

### Skema pengenal

| Jenis | Format | Contoh |
|---|---|---|
| Fitur | `F-<MOD>-NNN` | `F-FIN-004` |
| Kebutuhan fungsional | `FR-<MOD>-NNN` | `FR-REQ-012` |
| Business rule | `BR-<MOD>-NNN` | `BR-FIN-003` |
| Skenario | `SCN-NN` | `SCN-13` |
| Konsep inovasi | `IDEA-NN` | `IDEA-07` |
| Risiko | `RISK-NN` | `RISK-05` |
| Decision record | `ADR-NNN` | `ADR-002` |
| Pertanyaan terbuka | `Q-NN` | `Q-03` |
| Asumsi | `A-NN` | `A-07` |
| Riset yang perlu divalidasi | `R-NN` | `R-02` |
| Pertentangan antar-sumber | `K-NN` | `K-01` |
| Butir tinjauan hukum | `L-NN` | `L-05` |
| Contoh terhitung (dokumen 07) | `CTH-NN` | `CTH-21` |
| Invarian keuangan (dokumen 07) | `INV-NN` | `INV-06` |
| Invarian alur (dokumen 06) | `INV-ANN` | `INV-A04` |
| Pagar pencegah kendali (dokumen 06) | `PG-NN` | `PG-03` |
| Peristiwa notifikasi (dokumen 06) | `N-NN` | `N-02` |
| Butir terbuka dokumen 06 / 07 | `P-NN` / `T-NN` | `P-01`, `T-08` |
| Prinsip produk (dokumen 02) | `PRN-NN` | `PRN-06` |
| Ketegangan antar-prinsip (dokumen 02) | `KT-N` | `KT-9` |
| Persona (dokumen 03) | `PSN-NN` | `PSN-02` |
| Kebutuhan inti (dokumen 03) | `KB-NN` | `KB-12` |
| Masalah inti brief §1.1, kriteria sukses, pagar, tujuan (dokumen 02) | `MSL-N`, `KS-N`, `PAGAR`, `TUJUAN` | `MSL-5`, `KS-2` |
| Pertanyaan checkpoint, default, peringatan (dokumen 17) | `S2-N`, `D-NN`, `W-NN` | `S2-1`, `D-48`, `W-13` |
| Temuan pemeriksaan `security-privacy` | `SP-NN` | `SP-01` |
| Larangan visibilitas dan aturan tampilan (dokumen 06) | `V-N`, `AT-N` | `V-8`, `AT-4` |
| Transisi state machine (dokumen 06) | huruf + dua angka tanpa tanda hubung: `B`, `C`, `K`, `D`, `T`, `M`, `I` | `K06`, `C05` |

Baris sesudah `ADR-NNN` adalah tambahan atas skema di brief §16; tiap pengenal hanya dipakai di dokumen yang disebut. Pengenal transisi tidak bertanda hubung (`K06`) agar tidak tertukar dengan pengenal bertanda hubung berhuruf sama (`K-29` pertentangan, `T-08` butir terbuka, `D-48` default). `KB-NN` adalah kebutuhan inti dokumen 03; `KB` tanpa nomor adalah kolom kekayaan bersih di tabel dokumen 07.

Kode modul: `FAM`, `COM`, `FIN`, `REQ`, `SHP`, `INV`, `PLN`, `GOL`, `MNT`, `MEAL`, `CARE`, `CAL`, `TML`, `NTF`, `AI`. Kepanjangan tiap kode ada di [00-glossary.md](00-glossary.md).

### Kedalaman mengikuti tingkatan MVP

| Tingkatan | Kedalaman dokumentasi |
|---|---|
| Tier 1 — Inti | Kebutuhan lengkap dengan kriteria penerimaan |
| Tier 2 — Versi ringan | Kebutuhan ringkas |
| Tier 3 — Ditunda | Garis besar satu halaman per modul |

Cakupan tiap tingkatan didefinisikan di `family-os-brief.md` §1.1.

### Satu pemilik per dokumen

Setiap dokumen hanya diubah oleh pemiliknya. Permintaan perubahan atas dokumen milik pihak lain disampaikan lewat lead. Aturan uang, termasuk jumlah settlement, kanonis di dokumen 07; dokumen lain mengutipnya.
