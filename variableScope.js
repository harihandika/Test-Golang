// java script menganut konsep function scope dan BUKAN blok scope

"use strict" /*java script menggunakan mode strict yang mana di dalam function
hanya membuat variable lokal*/

// let a = 1;
// var b = 4;

// function tes() {
//     var b = 2 ;
//     console.log(a, " + " , window.b);
// }

// tes();
// console.log(a);

// var a = 1;
// function tes(a) {
//     console.log(a);
// }
// tes(2)

// function tampilAngka(n) {
//     if(n === 0) return;
//     console.log(n)
//     return tampilAngka(n-1)
// }
// tampilAngka(10);

// function faktorial (n) {
//     if (n === 0 ) return 1;
//     console.log(n)
//     return n * faktorial(n-1)
// }

// console.log(faktorial(5));

//  array
//  var myArr = ['teks',2,false, faktorial, [4,5,6]];
// console.log("ini arr", myArr[4][1]);

// Manipulasi Array
// 1. Menambah isi array
// var arr = ["a", 1, true];
// console.log(arr + " , " + arr[1] + ' , ' + arr[0]);

// var arr1 = [];
// arr[0] = "Hari";
// arr[1] = "Handika";
// arr[2] = "Setiawan";
// console.log(arr1);

// 2. Menghapus isi array
// var arr = ["Hari", "Handika", "Setiawan"];
// arr[1] = undefined;
// console.log(arr);

// 3. Menampilkan isi array
// var arr = ["Hari", "Handika", "Setiawan","Ganteng","Sekali"];

// for (var i = 0; i < arr.length; i++ ) {
//     console.log('Mahasiswa ke-' + (i+1) + ' : ' + arr[i]);
// }

// Method pada array
var arr = ["Hari","Handika","Setiawan","FullStack","Developer"];
// 1. Join
// console.log(arr.join(' - '));

// 2. push, pop
// arr.push('and Bakend','Developer')
// console.log(arr.join(' - '));
// arr.pop();
// console.log(arr.join(' - '));
// arr.pop();

// 3. shift, unshift
// arr.unshift('and Backend')
// arr.shift();
// console.log(arr.join(' - '));

// 4. splice
// splice(indexAwal, mauDihapusBerapa, elemenBaru1, elemenBaru2, ...)
// arr.splice(3, 0, 'and', 'Backend');
// arr.splice(2,2, 'Ini', 'Pengganti')
// console.log(arr.join(' - '));

// 5. slice
// slice(awal,akhir) note: dimana index akhir tidak dihitung
// var arr1 = arr.slice(1,4)
// console.log(arr1.join(' - '));

 //6. foreach
//  var angka = [1,2,3,4,5,6,7,8];
//  var nama = ['Hari', 'Handika', 'Setiawan', 'and', 'Backend']
//  for (var i = 0; i < angka.length; i++ ) {
//     console.log(angka[i]);
//  }
// var cetak = function(e) {
//     console.log(e);
// }
// angka.forEach(cetak);
// angka.forEach(function(e) {
//     console.log(e);
// })
// nama.forEach(function(e, i) {
//     console.log('Mahasiswa ke-' + (i+1) + ' adalah : ' + e);
// })

// 7. map
// var angka = [1,2,3,4,5,6,7,8];
// forEach tidak bisa dan jadi error
// var angka1 = angka.forEach(function(e) {
//     return e * 2;
// });
// console.log(angka1.join(' - '));

// var angka2 = angka.map(function(e) {
//     return e * 2;
// });
// console.log(angka2.join(' - '));

// 8 sort
// var angka = [1,2,8,10,20,6,5,7,4,3];
// console.log(angka.join(' - '));
// angka.sort(function(a,b){
//     return a-b
// });
// console.log(angka.join(' - '));

// 9. filter & find
// var angka = [1,2,3,4,5,10,8,6,9,7];
// var angka1 = angka.filter(function(x){
//     return x > 5;
// });
// console.log(angka1.sort(function(a,b){
//     return a-b
// }));
// var angka1 = angka.find(function(x){
//     return x > 5;
// });
// console.log(angka1);
