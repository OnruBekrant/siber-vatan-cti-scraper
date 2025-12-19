# 🛡️ Siber Vatan CTI Scraper

Bu proje, **Siber Vatan - Yıldız CTI Takımı** çalışmaları kapsamında geliştirilmiş bir Siber Tehdit İstihbaratı (CTI) veri toplama aracıdır.

## 🚀 Özellikler
- **Dinamik Hedefleme:** Hedef URL komut satırından verilir.
- **Stealth Mode (Gizlilik):** WAF ve Bot korumalarını (Cloudflare, Captcha vb.) aşmak için özel User-Agent ve tarayıcı bayrakları kullanır.
- **Full Page Screenshot:** Sayfanın tamamının ekran görüntüsünü alır.
- **Link Extraction:** Sayfadaki tüm bağlantıları analiz eder ve listeler.
- **Akıllı Hata Yönetimi:** Bağlantı hatalarını ve güvenlik duvarı (WAF) engellemelerini tespit edip raporlar.

## 🛠️ Kurulum

```bash
# Projeyi klonlayın
git clone [https://github.com/KULLANICI_ADIN/siber-vatan-cti-scraper.git](https://github.com/KULLANICI_ADIN/siber-vatan-cti-scraper.git)

# Gerekli bağımlılıkları yükleyin
go mod tidy
