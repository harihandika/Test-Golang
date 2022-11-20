// var penumpang = [];
// var penumpangNaik = function(namaPenumpang, penumpang) {
//     // jika angkot kosong 
//     if( penumpang.length == 0) {
//         //tambah penumpang di awal array
//         penumpang.push(namaPenumpang);
//         //kembalikan isi array & keluar dari function
//         return penumpang;
//     } else {
//         // telusuri seluruh kursi dari awal
//         for( var i = 0; i < penumpang.length; i++) {
//             // jika ada kursi kosong
//             if(penumpang[i] == undefined) {
//                // tambah penumpang di kursi tersebut
//                penumpang[i] = namaPenumpang;
//                //kembalikan isi array & keluar dari function
//                return penumpang;
//             }
//             // jika sudah ada nama yang sama
//             else if(penumpang[i] == namaPenumpang){
//             // tampilkan pesan kesalahannya
//             console.log(namaPenumpang + ' sudah ada di dalam angkot ^_^');
//             // kembalikan isi array & keluar dari function
//             return penumpang;
//             // jika seluruh kursi terisi
//             }
//             else if (i == penumpang.length - 1) {
//                 // tambah penumpang diakhir array
//                 penumpang.push(namaPenumpang);
//                 // kembalikan isi array & keluar dari function
//                 return penumpang;
//             }
//         }
        
//     }
// }

// var penumpangTurun = function(namaPenumpang, penumpang) {
//     if(penumpang.length == 0) {
//         console.log('Angkot masih kosong');
//         return penumpang;
//     } else {
//         for(var i = 0; i < penumpang.length; i++ ) {
//             if (penumpang[i] == namaPenumpang ) {
//                 penumpang[i] = undefined;
//                 return penumpang;
//             } else if( i == penumpang.length - 1 ) {
//                 console.log(namaPenumpang + ' tidak ada didalam angkot ');
//                 return penumpang;
//             }
//         }
//     }
// }


// Object dalam javascript
// var mhs = {
//     nama : 'Hari',
//     umur : 24,
//     ipsemester : [3.91, 3.80, 4.00, 3.97, 3.98, 3.87],
//     alamat : {
//         jalan : "Jl. Tegal Binangun",
//         kota : 'Palembang',
//         provinsi: 'Sumatera Selatan'
//     }
// };

// Membuat object
// object literal
// var mhs1 = {
//     nama : "Hari Handika Setiawan",
//     nrp : '030412712',
//     email : 'harihandika@gmail.com',
//     jurusan : 'teknik elektro'
// }

// var mhs2 = {
//     nama : "Michael Ouraahh",
//     nrp : '030412712',
//     email : 'ouraahh@gmail.com',
//     jurusan : 'teknik informatika'
// }

// // function declaration
// function buatObjectMahasiswa(nama, nrp, email, jurusan) {
//     var mhs = {};
//     mhs.nama = nama;
//     mhs.nrp = nrp;
//     mhs.email = email;
//     mhs.jurusan = jurusan;
//     return mhs;
// }

// var mhs3 = buatObjectMahasiswa('Jhon Cena', '03020302', 'jhon@gmail.com','teknik perkapalan')
// var mhs4 = buatObjectMahasiswa('Cynthia', '01020304','Cynthia@gmail.com', 'teknik informatika')

// // Constructor
// function Mahasiswa(nama , nrp, email, jurusan) {
//     // var this = {};
//     this.nama = nama;
//     this.nrp = nrp;
//     this.email = email;
//     this.jurusan = jurusan;
//     // return this;
// }

// var mhs5 = new Mahasiswa('Manda', '01020304', 'manda@gmail.com','Ekonomi')

// // cara 1 - function declaration
// function halo() {
//     console.log(this);
//     console.log('halo');
// }
// window.halo();
// // this mengembalikan object global

// // cara 2 - object literal
// var obj = {a : 10, nama : 'Hari'};
// obj.halo = function() {
//     console.log(this);
//     console.log('halo');
// }
// obj.halo()
// // this mengembalikan object yang bersangkutan

// // cara 3 - constructor
// function Halo() {
//     console.log(this);
//     console.log('halo');
// }
// var obj1 = new Halo();
// var obj2 = new Halo();
// // this mengembalikan object yang baru dibuat

// this
// console.log(window == this);
// window adalah object global


// var a = 10;
// console.log(this.a);
