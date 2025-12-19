#!/bin/bash

# Test edilecek 15 site listesi (CTI ve Haber siteleri karışık)
sites=(
    "https://www.sibervatan.org"
    "https://www.google.com"
    "https://www.haberler.com"
    "https://www.btk.gov.tr"
    "https://www.usom.gov.tr"
    "https://www.turkiye.gov.tr"
    "https://go.dev"
    "https://github.com"
    "https://stackoverflow.com"
    "https://www.hackread.com"
    "https://thehackernews.com"
    "https://www.darkreading.com"
    "https://portswigger.net"
    "https://tryhackme.com"
    "https://www.kali.org"
)

echo "🚀 Toplu Tarama Başlatılıyor..."

# Döngü ile her site için scraper'ı çalıştır
for site in "${sites[@]}"
do
   echo "------------------------------------------------"
   echo "📡 Taranıyor: $site"
   go run main.go "$site"
   echo "😴 Tarayıcıyı dinlendirmek için 2 saniye bekle..."
   sleep 2
done

echo "------------------------------------------------"
echo "✅ Tüm taramalar tamamlandı! 'outputs' klasörünü kontrol et."