package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// 1. Argüman Kontrolü
	if len(os.Args) < 2 {
		fmt.Println("❌ Hata: Lütfen bir URL giriniz.")
		os.Exit(1)
	}
	targetURL := os.Args[1]

	// 2. Klasör Hazırlığı
	safeURL := strings.ReplaceAll(targetURL, "https://", "")
	safeURL = strings.ReplaceAll(safeURL, "http://", "")
	safeURL = strings.ReplaceAll(safeURL, "/", "_")
	safeURL = strings.ReplaceAll(safeURL, ":", "")
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	
	// Klasör ismini biraz kısaltalım (okunabilirlik için)
	if len(safeURL) > 50 {
		safeURL = safeURL[:50]
	}
	
	folderName := fmt.Sprintf("%s_%s", safeURL, currentTime)
	outputDir := filepath.Join("outputs", folderName)

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		fmt.Printf("❌ Klasör hatası: %v\n", err)
		os.Exit(1)
	}

	// 3. Chromedp Ayarları (STEALTH MODU - GİZLİLİK) 🕵️‍♂️
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// A. User-Agent: Kendimizi normal bir Windows Chrome kullanıcısı gibi tanıtalım
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
		
		// B. Pencere Boyutu: Standart masaüstü
		chromedp.WindowSize(1920, 1080),
		
		// C. Otomasyon İzlerini Gizle (WAF Bypass için kritik!)
		chromedp.Flag("disable-blink-features", "AutomationControlled"), // "Ben robot değilim" bayrağı
		chromedp.Flag("enable-automation", false),                       // Otomasyon uyarısını kapat
		chromedp.NoFirstRun,                                             // İlk çalıştırma sihirbazlarını kapat
		chromedp.NoDefaultBrowserCheck,                                  // Varsayılan tarayıcı kontrolünü kapat
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Tarayıcıyı başlat
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Zaman aşımı (WAF'a takılırsak sonsuza kadar beklemeyelim)
	ctx, cancel = context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	var htmlContent string
	var screenshotBuffer []byte
	var links []string
	var title string

	fmt.Printf("🕵️‍♂️  Gizli Modda Bağlanılıyor: %s\n", targetURL)

	// 4. Görevler
	err := chromedp.Run(ctx,
		chromedp.Navigate(targetURL),
		// Sayfanın biraz yüklenmesini bekle (Cloudflare bazen 5sn bekletir)
		chromedp.Sleep(5*time.Second), 
		
		// Başlığı al (Hata analizi için)
		chromedp.Title(&title),

		// HTML'i Al
		chromedp.OuterHTML(`html`, &htmlContent, chromedp.ByQuery),

		// Linkleri Topla
		chromedp.Evaluate(`Array.from(document.querySelectorAll('a')).map(a => a.href)`, &links),

		// Tam Ekran Görüntüsü
		chromedp.FullScreenshot(&screenshotBuffer, 85),
	)

	// 5. AKILLI HATA ANALİZİ 🧠
	if err != nil {
		// Hata mesajını string'e çevirip analiz edelim
		errStr := err.Error()

		if strings.Contains(errStr, "deadline exceeded") {
			fmt.Println("⏳ HATA: Bağlantı Zaman Aşımına Uğradı!")
			fmt.Println("   👉 Sebep: Hedef site (WAF/Firewall) bağlantıyı engelliyor veya çok yavaşlatıyor.")
			fmt.Println("   👉 Durum: BTK/USOM gibi sitelerde bu durum normaldir (Güvenlik Önlemi).")
		} else {
			fmt.Printf("❌ Beklenmedik Hata: %v\n", err)
		}
		// Hata olsa bile klasörü temizlemeyelim, belki screenshot alınmıştır diyecektim ama
		// Run fonksiyonu hata verirse screenshot değişkeni boş kalır.
		// O yüzden burada çıkış yapıyoruz.
		return 
	}

	// 6. İÇERİK ANALİZİ (Cloudflare / Captcha Kontrolü)
	// HTML başarılı gelse bile içinde "Access Denied" yazıyor olabilir.
	securityKeywords := []string{"Cloudflare", "Captcha", "robot", "Access denied", "Attention Required"}
	detectedSecurity := false
	
	for _, keyword := range securityKeywords {
		if strings.Contains(title, keyword) || strings.Contains(htmlContent, keyword) {
			if !detectedSecurity { // Sadece bir kez yazdır
				fmt.Println("⚠️  UYARI: Bot Koruması Tespit Edildi!")
				detectedSecurity = true
			}
			fmt.Printf("   👉 Tespit edilen anahtar kelime: '%s'\n", keyword)
		}
	}
	
	if detectedSecurity {
		fmt.Println("   👉 Not: Ekran görüntüsü muhtemelen Captcha sayfasını gösterecektir.")
	}

	// 7. Dosyaları Kaydetme
	// HTML
	if err := os.WriteFile(filepath.Join(outputDir, "output.html"), []byte(htmlContent), 0644); err != nil {
		fmt.Println("❌ HTML kaydedilemedi.")
	} else {
		fmt.Printf("✅ HTML Kaydedildi (%d karakter)\n", len(htmlContent))
	}

	// Screenshot
	if len(screenshotBuffer) > 0 {
		if err := os.WriteFile(filepath.Join(outputDir, "screenshot.png"), screenshotBuffer, 0644); err != nil {
			fmt.Println("❌ Ekran görüntüsü kaydedilemedi.")
		} else {
			fmt.Printf("✅ Ekran Görüntüsü Kaydedildi (Boyut: %.2f KB)\n", float64(len(screenshotBuffer))/1024)
		}
	}

	// Links
	if len(links) > 0 {
		linksPath := filepath.Join(outputDir, "links.txt")
		linkData := strings.Join(links, "\n")
		if err := os.WriteFile(linksPath, []byte(linkData), 0644); err != nil {
			fmt.Println("❌ Linkler kaydedilemedi.")
		} else {
			fmt.Printf("✅ %d Adet Link Ayrıştırıldı ve Kaydedildi.\n", len(links))
		}
	} else {
		fmt.Println("🔸 Uyarı: Sayfada hiç link bulunamadı (veya korumaya takıldı).")
	}
	
	fmt.Println("🎉 İşlem Tamamlandı.\n")
}