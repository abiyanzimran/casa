# ADR-029 — Perangkat pengukuran kinerja ES §3

| | |
|---|---|
| Status | **Diterima**: diterima atas saran lead (keputusan 50, 12 November 2026); pemilik produk memilih "Setuju bunyi ini (Saran)" atas pratinjau yang disalin kata demi kata di [_progress.md](../_progress.md), "Penutupan pertanyaan terbuka sebelum S6". **ADR ini hanya mencakup kalimat ES §3.** Baris brief §1 "Pilot devices" dan "Platforms" disetujui dalam pratinjau yang sama dan dipasang lead di brief, bukan lewat ADR ini. Kalimat ES dipasang lead di `ENGINEERING-STANDARDS.md` §3. Butir 1 adalah tindakan lead yang dicatat di `_progress.md`; butir 2 mengulang kalimat yang disetujui; butir 3 dan 4 **[Usulan]** `architect`, tidak ada di pratinjau |
| Tanggal | 12 November 2026 (sesudah S5, sebelum S6; catatan web lebih dulu sesudah S6); 13 November 2026 bunyi ES §3 untuk web dipasang |
| Pengusul | `architect`, atas K-66 dan K-67 (17) |
| Mengubah | `ENGINEERING-STANDARDS.md` §3, paragraf pembuka (satu paragraf; ES baris 3: aturan ES hanya diubah lewat ADR yang disetujui pemilik produk) |
| Rincian | [14 §2.1, §4.10, §5.3](../14-engineering-roadmap.md) (QG-30 butir 4; QG-30 (Android) butir A-4); [ADR-012](ADR-012-ui-native-per-platform-dan-versi-os-minimum.md) catatan urutan platform; 15 RISK-54 |

## Konteks

- Bunyi ES §3 sebelum ADR ini: "Targets marked `[DEFAULT]` are starting values, confirmed or changed by ADR after the first measurement in the build phase on the pilot phones (iPhone 13, Galaxy A55). The iPad mini 7 is checked for layout, not as a performance target."
- Keputusan 42 (17): pilot dengan data nyata berjalan di iOS dan iPadOS, istri di iPhone 13 dan pemilik produk di iPad mini 7; GR-2 tidak menunggu Android; Galaxy A55 milik pemilik produk menyusul lewat jalur Android (14 §3.8); sebelum jalur pengujian Google Play dibuka (GR-4 butir 8; QG-31 (b)) (semula bawaan lead; diterima atas saran lead lewat keputusan 50, dengan syarat beta dari keputusan 51; bunyi brief §1 "Android is built after real data starts and joins the pilot before the beta, if a beta is ever held").
- Akibatnya dua hal di bunyi lama tidak cocok lagi (17 K-67): (1) iPad mini 7 adalah perangkat harian salah satu pasangan, jadi kinerjanya ikut menentukan usaha mencatat (KS-1; RISK-19), bukan hanya tata letaknya; (2) Galaxy A55 belum ada di FS-1, jadi pengukuran pertama di sana tidak dapat menjadi syarat data nyata pertama.
- 14 dan 13 sudah memakai urutan yang baru (14 §2.1, QG-30 butir 4, QG-30 (Android) butir A-4; 13 §3.3 butir 5), sambil menunggu ES diubah (K-67: keputusan menang sampai ES diubah). ADR ini menutup selisih itu.
- Sasaran dan angkanya sendiri tidak dipersoalkan. Yang berubah hanya di perangkat mana dan kapan angka `[DEFAULT]` dikukuhkan.

## Opsi

| Kriteria | A. Bunyi lama tetap; selisih dicatat di 17 | **B. Perangkat pengukuran menurut urutan platform pilot** | C. Ukur ketiga perangkat sebelum GR-2 |
|---|---|---|---|
| Sesuai keputusan 42 | Tidak: Galaxy A55 menjadi syarat yang tidak dapat dipenuhi sebelum GR-2 | Ya | Tidak: GR-2 menunggu Android |
| iPad mini 7 diukur sebagai perangkat harian | Tidak | Ya | Ya |
| Galaxy A55 tetap diukur sebelum memegang data nyata | Ya, tetapi terlalu awal | Ya, sebelum aplikasi Android memegang data nyata | Ya |
| Pekerjaan tambahan | Tidak ada; ES dan rencana tetap berbeda | Tidak ada; rencana sudah seperti ini | Seluruh klien Android sebelum GR-2 |

## Keputusan

**Opsi B.** Paragraf pembuka ES §3 diganti dengan kalimat berikut, sebagaimana terpasang di ES §3 baris 29:

> Targets marked `[DEFAULT]` are starting values, confirmed or changed by ADR after the first measurement in the build phase on the pilot devices: iPhone 13 and iPad mini 7 before the first real data, Galaxy A55 before the Android app holds real data (ADR-029).

Bunyi yang disetujui pemilik produk adalah kalimat itu tanpa "(ADR-029)"; rujukan dalam kurung ditambahkan lead saat memasangnya, mengikuti pola rujukan ADR lain di ES, dan tidak mengubah isi.

1. Kalimat lama "The iPad mini 7 is checked for layout, not as a performance target" **dihapus** lead saat memasang, karena bertentangan dengan kalimat baru dan tidak ada di pratinjau (`_progress.md`). Pemeriksaan tata letak iPad tetap dituntut di tempat lain (08 §16.4; 14 QG-30 butir 3).
2. **iPhone 13 dan iPad mini 7:** pengukuran pertama dan ADR yang mengukuhkan atau mengubah angka `[DEFAULT]` selesai sebelum GR-2 (14 QG-30 butir 4).
3. **Galaxy A55:** diukur sebelum APK pertama dipasang di perangkat yang memegang data nyata (14 QG-30 (Android) butir A-4). **[Usulan]** `architect`: bila angkanya berbeda, ADR kinerja yang sama direvisi, bukan ADR kedua (14 §4.10).
4. Tabel sasaran ES §3 dan butir di bawahnya tidak berubah. **[Usulan]** `architect`, bacaan: "pilot devices" di baris "Scrolling and transitions" berarti ketiga perangkat, masing-masing sejak ia memegang data nyata.

## Konsekuensi

- ES, 13, dan 14 berbunyi sama; K-66 dan K-67 di 17 ditutup dengan keputusan 50.
- Sampai jalur Android, angka `[DEFAULT]` dikukuhkan tanpa perangkat Android. Bila Galaxy A55 ternyata perangkat paling lambat **[Asumsi]** (RISK-54), angka yang dikukuhkan di GR-2 dapat direvisi di QG-30 (Android); data nyata di iOS tidak menunggu revisi itu.
- iPad mini 7 kini perangkat ukur: cold start dan kecepatan alur Catat diukur juga di iPad sebelum GR-2. Pekerjaannya sudah ada di 14 §3.3 langkah 7 dan §4.10; tidak ada pekerjaan baru.
- Tidak ada biaya uang.

## Catatan keputusan web lebih dulu (12 November 2026, sesudah S6)

Mencatat [ADR-030](ADR-030-klien-web-pwa-lebih-dulu.md) butir 16. **Bunyi ES §3 baru (pengukuran di klien web; baris ukuran app shell web ≤ 500 KB `[DEFAULT]`) disetujui pemilik produk dan dipasang lead (13 November 2026)** (`_progress.md`, "Persetujuan pemilik produk atas bunyi brief, `CLAUDE.md`, dan ES").

- **Sebelum GR-2:** sasaran `[DEFAULT]` ES §3 diukur di **klien web** pada ketiga perangkat pilot:
  - iPhone 13 dan iPad mini 7, web app layar utama, Safari 26;
  - Galaxy A55, Chrome.

  Galaxy A55 memegang data nyata sejak GR-2 lewat web, jadi alasan butir 3 (menunggu jalur Android) tidak berlaku bagi web. Ini kini bunyi ES §3 yang berlaku.
- **Padanan web untuk sasaran ES §3** **[Usulan]**:

  | Sasaran ES | Diukur di web sebagai |
  |---|---|
  | Cold start ≤ 2 detik | Dari ketukan ikon web app sampai Beranda dengan data lokal, app shell dari cache service worker |
  | "Saved locally at once" | Tulis IndexedDB dan isian tertutup |
  | Scrolling and transitions | Tanpa frame terbuang yang terlihat, diukur di Chrome DevTools jarak jauh untuk Galaxy A55 dan dengan pengamatan di iPhone dan iPad (tanpa Web Inspector; ADR-022 catatan web) |
  | App download size ≤ 30 MB | Ukuran transfer app shell terkompresi ≤ 500 KB **[Asumsi]**, dikukuhkan sesudah prototipe |
- **Aplikasi native:** butir 2 dan 3 tetap berlaku untuk masing-masing saat ia akan memegang data nyata. ADR kinerja yang sama direvisi, bukan ADR kedua.
- **Bunyi ES §3 yang diusulkan** (ADR-030 butir 16; disetujui pemilik produk dan dipasang lead, 13 November 2026, dengan rumusan lead; bunyi persisnya di `ENGINEERING-STANDARDS.md` §3): "…on the pilot devices: the web client on iPhone 13, iPad mini 7 and Galaxy A55 before the first real data; each native app on its pilot devices before it holds real data (ADR-029, ADR-030)."
