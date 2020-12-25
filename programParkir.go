/*KELOMPOK 11
	KELAS : IF-43-05
	-> NAMA : AZIZAH CAHYA KEMILA
		NIM : 1301194103
	-> NAMA : TSAQIB SAYYIDAN SENDJAJA
		NIM : 1301193450
*/
	
package main

import "fmt"
import "strings"
import "os"
import "os/exec"

const T = 10
const N = 5
const M = 5

type parkir struct {
	noPolisi       string
	jenisKendaraan string
	tanggal        int
	bulan          int
	tahun          int
	jam            int
	menit          int
}

type arrayMotor [T]parkir
type arrayMobil [N][M]parkir


var reset = parkir{
	noPolisi:       "",
	jenisKendaraan: "",
	jam:            0,
	menit:          0,
	tanggal:        0,
	bulan:          0,
}

var data arrayMotor
var data2 arrayMobil
var nMobil, nMotor, nBus, totalTarif, cMobil,cMotor int

func menu(){
	fmt.Println("================= PARKING TICKET =================")
	fmt.Println()
	fmt.Println("TRANSMART PODOMORO DAYEUHKOLOT")
	fmt.Println("  *Parkir buka dari jam 09.00 WIB - 22.00 WIB*")
	fmt.Println()
	fmt.Println("MENU")
	fmt.Println("1. Masukkan data dan lokasi parkir kendaraan")
	fmt.Println("2. Merapikan parkir motor ")
	fmt.Println("3. Mencari kendaraan (No. Polisi)")
	fmt.Println("4. Mencetak kartu keluar parkir")
	fmt.Println("5. Mencetak No.Polisi kendaraan berdasarkan jenisnya")
	fmt.Println("6. Statistik harian")
	fmt.Println("7. Pendapatan sehari")
	fmt.Println("8. exit")
}
func main() {
	var pilih int
	var press string
	menu()
	fmt.Print("PILIH MENU: ")
	fmt.Scanln(&pilih)

	for pilih < 8 {

		switch {
		case pilih == 1:
			clear()	
			isiParkir()
		case pilih == 2:
			clear()
			merapikanParkir()
		case pilih == 3:
			clear()
			mencariKendaraan()
		case pilih == 4:
			clear()
			tarifKeluar()
		case pilih == 5:
			clear()
			sort()
		case pilih == 6:
			clear()
			jumlahKendaraan()
			okupansi()
		case pilih == 7:
			clear()
			pendapatan()
		}
		fmt.Print("INPUT SESUATU UNTUK MELANJUTKAN: ")
		fmt.Scanln(&press)
		clear()
		menu()
		fmt.Print("PILIH MENU: ")
		fmt.Scanln(&pilih)

	}
}

func isiParkir() {
	/* I.S 
		F.S array terisi dengan data kendaraan*/
	var lokasi, baris, kolom int
	var kendaraan parkir
	fmt.Println("MASUKKAN DATA KENDARAAN")
	fmt.Print("Nomor Polisi: ")
	fmt.Scanln(&kendaraan.noPolisi)
	fmt.Print("Jenis Kendaraan (motor, mobil, atau bus): ")
	fmt.Scanln(&kendaraan.jenisKendaraan)
	kendaraan.jenisKendaraan = strings.ToLower(kendaraan.jenisKendaraan)
	for kendaraan.jenisKendaraan != "motor" && kendaraan.jenisKendaraan != "mobil" && kendaraan.jenisKendaraan != "bus" {
		fmt.Println("Maaf jenis kendaraan yang anda masukkan tidak valid")
		fmt.Print("Masukkan jenis kendaraan (motor, mobil, atau bus): ")
		fmt.Scanln(&kendaraan.jenisKendaraan)
		kendaraan.jenisKendaraan = strings.ToLower(kendaraan.jenisKendaraan)
	}
	fmt.Print("Tanggal Masuk(DD MM YY): ")
	fmt.Scanln(&kendaraan.tanggal, &kendaraan.bulan, &kendaraan.tahun)
	for kendaraan.tanggal > 31 || kendaraan.tanggal <= 0 || kendaraan.bulan > 12 || kendaraan.bulan <= 0 || (kendaraan.bulan > 28) {
		fmt.Println("Maaf tanggal yang anda masukkan tidak valid")
		fmt.Print("Tanggal Masuk(DD MM YY): ")
		fmt.Scanln(&kendaraan.tanggal, &kendaraan.bulan, &kendaraan.tahun)
	}
	fmt.Print("Waktu Masuk(hh mm): ")
	fmt.Scanln(&kendaraan.jam, &kendaraan.menit)
	for kendaraan.jam > 22 || kendaraan.jam < 9 || kendaraan.menit >=60 || kendaraan.menit < 0{
		fmt.Println("Maaf waktu yang anda masukkan tidak valid")
		fmt.Print("Waktu Masuk(hh mm): ")
		fmt.Scanln(&kendaraan.jam, &kendaraan.menit)
	}
  
	if kendaraan.jenisKendaraan == "motor" {
		fmt.Print("Masukkan lokasi parkir(contoh : 1) : ")
		fmt.Scanln(&lokasi)
		
		for lokasi < 0 || lokasi > T-1 || data[lokasi].noPolisi != "" {
			if 	lokasi < 0 || lokasi > T-1{
				fmt.Println("MAAF LOKASI PARKIR TIDAK VALID")
				fmt.Print("Silahkan masukkan lokasi parkir (contoh : 1) : ")
				fmt.Scanln(&lokasi)
			}else if data[lokasi].noPolisi != ""{
				fmt.Println("PARKIR SUDAH TERISI")
				fmt.Print("Silahkan masukkan lokasi parkir (contoh : 1) : ")
				fmt.Scanln(&lokasi)
			}
		}
		data[lokasi] = kendaraan
		nMotor = nMotor + 1
		cMotor++
	} else if kendaraan.jenisKendaraan == "mobil" {
		fmt.Print("Masukkan lokasi parkir (contoh : 1 2) : ")
		fmt.Scanln(&baris, &kolom)
		for baris < 0 || kolom < 0 || baris > N-1 || kolom > M-1 || data2[baris][kolom].noPolisi != ""{
			if baris < 0 || kolom < 0 || baris > N-1 || kolom > M-1{
				fmt.Println("MAAF LOKASI PARKIR TIDAK VALID")
				fmt.Print("Silahkan masukkan lokasi parkir (contoh : 1 2) : ")
				fmt.Scanln(&baris, &kolom)
			}else if data2[baris][kolom].noPolisi != "" {
				fmt.Println("MAAF PARKIR SUDAH TERISI")
				fmt.Print("silahkan masukkan lokasi parkir (contoh : 1 2) : ")
				fmt.Scanln(&baris, &kolom)
			} 
		}
		data2[baris][kolom] = kendaraan
		nMobil = nMobil + 1
		cMobil++
	} else if kendaraan.jenisKendaraan == "bus" {
		fmt.Print("Masukkan lokasi parkir (contoh: 1 2): ")
		fmt.Scanln(&baris, &kolom)
		for baris < 0 || kolom < 0 || baris > N-1 || kolom > M-1 || data2[baris][kolom].noPolisi != "" {
			if baris < 0 || kolom < 0 || baris > N-1 || kolom > M-1 {
				fmt.Println("MAAF LOKASI PARKIR TIDAK VALID")
				fmt.Print("Silahkan masukkan lokasi parkir (contoh : 1 2) : ")
				fmt.Scanln(&baris, &kolom)
			}else if data2[baris][kolom+1].noPolisi != "" {
				fmt.Println("MAAF PARKIR SUDAH TERISI")
				fmt.Print("Masukkan lokasi parkir (contoh : 1 2) : ")
				fmt.Scanln(&baris, &kolom)
			}
		}
		if (baris+1) > N-1 && (kolom+1) > M-1  {
			fmt.Print("Maaf lokasi parkir tidak valid")
		}else if kolom + 1 > M-1   {
			kolom = 0
			fmt.Println("Mohon maaf lokasi parkir tidak muat untuk bus, jadi dipindahkan ke ", baris+1, kolom)
			data2[baris+1][kolom] = kendaraan
			data2[baris+1][kolom+1] = kendaraan
		}else {
			data2[baris][kolom] = kendaraan
			data2[baris][kolom+1] = kendaraan
		}
		nBus = nBus + 1
		cMobil++
	}
	fmt.Println()
	fmt.Println("================= PARKING TICKET =================")
	fmt.Println()
	fmt.Println("DATA KENDARAAN")
	fmt.Println("Nomor Kendaraan: ", kendaraan.noPolisi)
	fmt.Println("Jenis Kendaraan (motor, mobil, atau bus): ", kendaraan.jenisKendaraan)
	fmt.Println("Tanggal (DD MM YY): ", kendaraan.tanggal, kendaraan.bulan, kendaraan.tahun)
	fmt.Println("waktu (hh mm): ", kendaraan.jam, kendaraan.menit)
	fmt.Println()
	fmt.Println("==================================================")
}

func tarifKeluar() {
	/* I.S durasi, tarif, jamKeluar, menitKeluar diinput oleh user 
		F.S mengeluarkan tarif parkir , mencatat waktu parkir, dan menghapus data kendaraan yang keluar, dari array */
	var durasi, tarif, jamKeluar, menitKeluar, i, j int
	var keluar, masuk, kendaraan parkir
	var nomor string
	var jenis string

	found := false
	fmt.Print("masukkan tarif parkir: ")
	fmt.Scanln(&tarif)
	fmt.Print("Masukkan nomor kendaraan: ")
	fmt.Scanln(&nomor)
	fmt.Print("Masukkan jenis kendaraan (motor, mobil, atau bus): ")
	fmt.Scanln(&jenis)
	fmt.Print("Keluar : ")
	fmt.Scanln(&keluar.jam, &keluar.menit)
	if jenis == "motor" {
		for i := 0; i < T; i++ {
			if nomor == data[i].noPolisi {
				found = true
				masuk = data[i]
			}
		}
	} else if jenis == "mobil" || jenis == "bus" {
		for j := 0; j < N; j++ {
			for k := 0; k < M; k++ {
				if nomor == data2[j][k].noPolisi {
					found = true
					masuk = data2[j][k]
				}
			}
		}
	}
	if found == true && jenis == "motor" {
		jamKeluar = (keluar.jam * 60) - (masuk.jam * 60)
		menitKeluar = keluar.menit - masuk.menit
		cetakTiket(masuk.noPolisi, masuk.jam, masuk.menit, keluar.jam, keluar.menit)
	} else if found == true && (jenis == "mobil" || jenis == "bus") {
		jamKeluar = (keluar.jam * 60) - (masuk.jam * 60)
		menitKeluar = keluar.menit - masuk.menit
		cetakTiket(masuk.noPolisi, masuk.jam, masuk.menit, keluar.jam, keluar.menit)
	}
	durasi = jamKeluar + menitKeluar
	jam := durasi / 60
	menit := durasi % 60
	if menit <= 10 {
		tarif = tarif * jam
	} else {
		tarif = tarif * (jam + 1)
	}
	fmt.Println("Durasi : ", jam, menit)
	fmt.Println("Total : ", tarif)
	totalTarif = totalTarif + tarif

	cariDataKendaraan(nomor, &kendaraan, &i, &j)

	if kendaraan.jenisKendaraan == "motor"{
		data[i] = reset
		cMotor = cMotor - 1
	}else if kendaraan.jenisKendaraan == "mobil"{
		data2[i][j] = reset
		cMobil = cMobil - 1
	}else if kendaraan.jenisKendaraan == "bus" {
		data2[i][j] = reset
		data2[i][j+1] = reset
		cMobil = cMobil - 1
	}

	fmt.Println("=============TERIMA KASIH TELAH BERKUNJUNG===========")
	
}

func merapikanParkir() {
	/*I.S array tidak rapi yaitu terdapat data kosong diantara data yang terisi sehingga perlu dirapikan
		F.S data telah rapi dan semua array kosong berada di sebelah kanan */
	var j int
		for i:= 0; i < T-1; i++ {
			j = i+1
			for j < T{
				if data[i].noPolisi == "" && data[j].noPolisi != "" {
					data[i] = data[j]
					data[j] = reset
				}
				j++
			}
		}
	fmt.Println(data)
}

func cariDataKendaraan(nomor string, kendaraan *parkir, x *int, y *int) {
	/* I.S menginputkan nomor, kendaraan, x, y
		F.S mengirimkan true jika kendaraan yang dicari ada pada array, 
			atau false jika kendaraan yang dicari tidak ada pada array */
	*kendaraan = reset
	*x = -1
	*y = -1

	found := false

	for i := 0; i < T && !found; i++ {
		if nomor == data[i].noPolisi {
			*kendaraan = data[i]
			*x = i

			found = true
		}
	
	}

	for j := 0; j < N; j++ {
		for k := 0; k < M && !found; k++ {
			if nomor == data2[j][k].noPolisi {
				*kendaraan = data2[j][k]
				*x = j
				*y = k

				found = true
			}
		}
	}
}

func mencariKendaraan() {
	/*I.S nomor, kendaraan sebagai inputan user
		F.S kendaraan ditemukan atau tidak */
	var nomor string
	var kendaraan parkir
	var i, j int

	fmt.Print("Masukkan nomor kendaraan: ")
	fmt.Scanln(&nomor)

	cariDataKendaraan(nomor, &kendaraan, &i, &j)

	if kendaraan.noPolisi == "" {
		fmt.Println("Maaf kendaraan yang anda cari tidak ditemukan")
	} else {
		fmt.Println("Kendaraan ditemukan")
		fmt.Println("Data Kendaraan: ", kendaraan)
	}
}

func sort() { 
	/* I.S arrMotor berfungsi untuk menampung arrayMotor sehingga data pada arrayMotor tidak berubah
			arrMobil untuk menampung arrayMobil dalam bentuk array 1 dimensi
			arrBus untuk menampung data bus, newArrMobil untuk menampung data mobil sehingga data pada arrMobil tidak berubah
		F.S mengeluarkan array masing-masing kendaraan yang terurut mengecil berdasarkan no.Polisi */
	var arrMotor arrayMotor
	var arrBus [N*M] parkir
	var newArrMobil [N*M] parkir
	var arrMobil [N*M] parkir
	var j, k, l, max int
	var t parkir
	
	arrMotor = data

	for p:=0; p < N; p++ {
		for q:= 0; q < M; q++ {
			arrMobil[p*N + q] = data2[p][q]
		}
	
	}
	for i := 0; i < (N*M); i++ {
		max = i
		j = i + 1
		for j < (N*M) {
			if arrMobil[j].noPolisi > arrMobil[max].noPolisi {
				max = j
			}
			j++
		}
		t = arrMobil[i]
		arrMobil[i] = arrMobil[max]
		arrMobil[max] = t
	}
	for i:=0;i< (N*M);i++{
		
		if arrMobil[i].jenisKendaraan == "mobil"{
			newArrMobil[k] = arrMobil[i]
			k++
		}else if arrMobil[i].jenisKendaraan == "bus"{
			arrBus[l] = arrMobil[i]
			l++
		}
	}
	fmt.Println("Parkir Mobil: ")
	for i := 0;i<(N*M);i++ {
		fmt.Print(newArrMobil[i].noPolisi, " ")
	}
	
	fmt.Println()
	fmt.Println("Parkir Bus: ")
	for i:=0;i<(N*M);i++{
		fmt.Print(arrBus[i].noPolisi," ")
	}
	fmt.Println()
	for i := 0; i < T-1; i++ {
		max = i
		j = i + 1
		for j < T {
			if arrMotor[j].noPolisi > arrMotor[max].noPolisi {
				max = j
			}
			j++
		}
		t = arrMotor[i]
		arrMotor[i] = arrMotor[max]
		arrMotor[max] = t
	}
	fmt.Println("Parkir Motor: ")
	for i:=0;i<T;i++{
		fmt.Print(arrMotor[i].noPolisi," ")	
	}
	fmt.Println()
	

}

func jumlahKendaraan(){
	/* I.S nMotor, nMobil digunakan untuk menghitung jumlah kendaraan yang telah parkir
		F.S mengeluarkan jumlah kendaraan berdasarkan jenisnya*/
	fmt.Println("JUMLAH KENDARAAN YANG TELAH PARKIR HARI INI BERDASARKAN JENISNYA")
	fmt.Println("Motor: ", nMotor)
	fmt.Println("Mobil: ", nMobil)
	fmt.Println("Bus: ", nBus)
}

func okupansi(){
	/*I.S cMotor, cMobil adalh total mobil dan motor yang berada diparkiran saat ini
		F.S output presentase okupansi kendaraan berdasarkan parkir motor dan mobil */
	var okupansiMotor, okupansiMobil float64

	okupansiMotor = float64(cMotor)/float64(T)*100
	okupansiMobil = float64(cMobil)/(float64(N)*float64(M))*100
	fmt.Println("PRESENTASI OKUPANSI PARKIR SAAT INI")
	fmt.Println("Area motor: ", okupansiMotor, "%")
	fmt.Println("Area Mobil: ", okupansiMobil, "%")
}

func pendapatan(){
	/* I.S totalTarif berisi total tarif parkir kendaraan yang keluar hari ini
		F.S  mengeluarkan pendapatan hari ini*/
	fmt.Println("Total pendapatan hari ini: ", totalTarif)
}


func clear(){
	a:=exec.Command("cmd","/c","cls")
	a.Stdout=os.Stdout
	a.Run()
}

func cetakTiket(noPolisi string, jam, menit, jamKeluar, menitKeluar int) {
	f, err := os.OpenFile("Tiket Parkir.txt", os.O_APPEND|os.O_CREATE, 0600) //if FOUND = append file. if !FOUND create file.

	if err != nil {
		fmt.Print("error creating file: ", err)
		return
	}
	defer f.Close()
	for i := 0; i < 1; i++ { // Generating...
		_, err = f.WriteString(fmt.Sprintf("nomor kendaraan = %s\njam masuk kendaraan = %02d : %02d\njam keluar kendaraan = %02d : %02d \n\n", noPolisi, jam, menit, jamKeluar, menitKeluar)) // writing...
		if err != nil {
			fmt.Print("error writing string: ", err)
		}
	}
}
