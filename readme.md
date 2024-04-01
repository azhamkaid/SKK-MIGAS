aplikasi berikut adalah berupa fungsi yang berguna untuk menghintung jumlah kemunculan masing-masing huruf dalam suatu kalimat dengan jumlah 5 kata

"Pendanaan", "Terproteksi", "Untuk", "Dampak", "Berarti".

proses dilakukan dengan mengabungkan kelima kata menjadi string tunggal menggunakan build in func strings.join

dengan hasil
PendanaanTerproteksiUntukDampakBerarti

selanjutnya melakukan looping untuk menhitung jumlah kemuncuman masing masing huruf dengan menggunakan angka rune dari A=65 hingga Z=90 dimana angka rune ini merupakan representasikan unicode dari setiap karakter menjadi sebuah int, sehingga kita dapat menghitung seberapa banyak kemunculan setiap huruf dalam kata "PendanaanTerproteksiUntukDampakBerarti", sehingga didapatkan [B:1 D:1 P:1 T:1 U:1 a:6 d:1 e:4 i:2 k:3 m:1 n:4 o:1 p:2 r:4 s:1 t:3 u:1] data ini di masukan dalam mapping (value := make(map[string]int)) selanjutnya mapping dengan data berisikan huruf dan jumlah kemunculan nya akan dilakukan sorting dengan syarat sebagai berikut

1. urutkan berdasarkan jumlah kemunculan paling banyak hingga paling sedikit
2. jika nilai nya sama prioritaskan huruf besar terlebih dahulu
3. jika nilai sama dan sama sama huruf besar, urutkan berdasarkan abjad
sehingga di dapatkan hasil 
[a e n r k t i p B D P T U d m o s u]






#   S K K - M I G A S  
 