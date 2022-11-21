//dengan deklarasi
// function jumlah (a,b){
//     let hasil;
//     hasil = a + b;

//     return hasil;
// }

//atau dengan ekspresi
// let jumlahExpres = function(a, b) {
//    let hasil;
//     hasil = a + b;

//     return hasil;
// }

// alert(jumlahExpres(1,5))
// alert(jumlahExpres(10,20))
// alert(jumlahExpres(20,30))

// var a = 10;
// var b = 3;

// volumeA = a * a * a
// volumeB = b * b * b
// total = volumeA + volumeB
// total1 = volumeA
// console.log(total, total1);

//  function volumeKubus(a, b) { 
//     let kubusA = a*a*a
//     let kubusB = b*b*b
//     let hasil = kubusA + kubusB
//     return hasil
// }
// console.log("ini function", volumeKubus(10,3))

// function tambah(a,b) {
//     return a + b;
// }

// function kali(a ,b) {
//      return a * b
// }
// let a =  tambah(1,2)
// let b = tambah(3,4)
// let hasil =  kali(a, b);
// let a = parseInt(prompt('Masukkan nilai 1: '));
// let b = parseInt(prompt('Masukkan nilai 2: '));
// let hasil = tambah(a+10,b+10)
// console.log(a,b);
// console.log(hasil);

//  function  tambah() {
// return arguments;
//  }

// let coba = tambah(5,10,20,"hi",false);
// console.log(coba);

function tambah() {
    let hasil = 0;
    for (let i = 0; i < arguments.length; i++) {
        hasil += arguments[i];
    }
    return hasil;
}

let coba = tambah(1,2,3,4,5);
console.log(coba);
