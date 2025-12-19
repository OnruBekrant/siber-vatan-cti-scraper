# 🛡️ Siber Vatan CTI Scraper

> **Siber Vatan - Yıldız CTI Takımı** çalışmaları kapsamında geliştirilmiş, yeni nesil Siber Tehdit İstihbaratı (CTI) veri toplama aracıdır.

Bu araç, hedef web sitelerinden **otonom** bir şekilde veri toplamak, görsel kanıt (screenshot) almak ve bağlantı analizi yapmak için tasarlanmıştır. Özellikle **WAF (Web Application Firewall)** ve **Anti-Bot** sistemlerini atlatmaya yönelik gelişmiş "Gizlilik Modu" (Stealth Mode) özelliklerine sahiptir.

## 🚀 Özellikler

* **🕵️‍♂️ Gizlilik Modu (Stealth Mode):** Cloudflare, Captcha ve gelişmiş bot korumalarını aşmak için özel `User-Agent` manipülasyonu ve otomasyon bayraklarını gizleme teknikleri kullanır.
* **📸 Tam Sayfa Ekran Görüntüsü (Full Page Screenshot):** Sadece görünen alanı değil, sayfanın tamamını (scroll ederek) yüksek kalitede kaydeder.
* **🔗 Bağlantı Analizi (Link Extraction):** Hedef sayfadaki tüm iç ve dış bağlantıları (`href`) analiz eder ve listeler.
* **🧠 Akıllı Hata Yönetimi:** Bağlantı hatalarını analiz eder; sunucu hatası mı yoksa güvenlik duvarı (WAF) engellemesi mi olduğunu tespit edip raporlar.
* **📂 Dinamik Kayıt Sistemi:** Her tarama için `outputs/DOMAIN_TARIH` formatında benzersiz klasörler oluşturarak verileri düzenli tutar.

## 🛠️ Kurulum

Projeyi yerel makinenize kurmak için aşağıdaki adımları izleyin:

### Gereksinimler
* [Go (Golang)](https://go.dev/dl/) 1.20 veya üzeri
* Google Chrome veya Chromium Tarayıcı

### Adım Adım Kurulum

1. **Projeyi Klonlayın:**
    ```bash
    git clone [https://github.com/OnruBekrant/siber-vatan-cti-scraper.git](https://github.com/OnruBekrant/siber-vatan-cti-scraper.git)
    cd siber-vatan-cti-scraper
    ```

2. **Bağımlılıkları Yükleyin:**
    ```bash
    go mod tidy
    ```

## 💻 Kullanım

Aracı iki farklı modda kullanabilirsiniz:

### 1. Tekil Tarama (Single Scan)
Belirli bir hedefi taramak için URL'yi parametre olarak verin:
###

```bash
go run main.go [https://www.hedefsite.com](https://www.hedefsite.com)

2. **Toplu Tarama (Batch Scan / Otomasyon):**
    Birden fazla hedefi (liste halinde) otomatik taramak için hazırlanan scripti kullanın:
    ```bash
    chmod +x run_tests.sh
    ./run_tests.sh
    ```

## 📂 Proje Yapısı

```text
siber-vatan-cti-scraper/
├── 📂 outputs/          # Tarama sonuçlarının kaydedildiği dizin
│   └── 📂 site.com_.../ # Her siteye özel oluşturulan klasör
│       ├── 📄 output.html      # Sitenin kaynak kodları
│       ├── 🖼️ screenshot.png   # Tam sayfa ekran görüntüsü
│       └── 📄 links.txt        # Çıkarılan linklerin listesi
├── 📄 main.go           # Ana kaynak kod (Scraper motoru)
├── 📜 run_tests.sh      # Toplu tarama otomasyon scripti
├── 📄 go.mod            # Go modül dosyası
└── 📝 README.md         # Proje dokümantasyonu
⚠️ Yasal Uyarı

Bu araç, Siber Vatan eğitim programı kapsamında eğitim ve savunma amaçlı geliştirilmiştir. Hedef sistemlerin izni olmadan saldırı veya yetkisiz veri toplama amacıyla kullanılması yasaktır. Geliştirici, aracın kötüye kullanımından doğacak sonuçlardan sorumlu tutulamaz.

Geliştirici: Onur Berkant Girgeç

Tarih: 20.12.2025