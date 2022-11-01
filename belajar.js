// let jumlahAngkot = 10;
// var nilai = 1;
// while (nilai <= jumlahAngkot) {
//   console.log("Angkot No. " + nilai + " beroperasi dengan baik");
//   nilai++;
// }
// let jumlahAngkot = 10;
// let benar = 7;
// for (let nilai = 1; nilai <= jumlahAngkot; nilai++) {
//   if (nilai < benar) {
//     alert("hello " + nilai);
//   } else {
//     alert("Angkot " + nilai + "jelek");
//   }
// }

// let angka = prompt("Masukkan Angka");
// if (angka % 2 === 0) {
//   alert(angka + " ini adalah genap");
// } else if (angka % 2 == 1) {
//   alert(angka + "ini ganjil");
// } else {
//   alert("ini bukan angka");
// }

// let jumlahAngkot = 10;
// let lembur = 8;
// let beroperasi = 6;
// let angkot = 1;
// for (angkot = 1; angkot <= jumlahAngkot; angkot++) {
//   if (angkot <= beroperasi && angkot !== 5) {
//     console.log(" angkot no " + angkot + " beroperasi dengan baik");
//   } else if (angkot === lembur || angkot === 9 || angkot === 5) {
//     console.log("Angkot no " + angkot + " lembur");
//   } else {
//     console.log("Angkot No " + angkot + "Tidak beroperasi");
//   }
// }

// let angka = parseInt(prompt("masukkan angka : "));

// if (angka === 1) {
//   alert("angka ini 1");
// } else if (angka === 2) {
//   alert("angka ini 2");
// } else {
//   alert("angka salah");
// }

// let angka = prompt("masukkan angka :");

// switch (angka) {
//   case "1":
//     alert("anda memasukkan angka 1");
//     break;
//   case "2":
//     alert("anda memasukkan angka 2");
//     break;
//   case "3":
//     alert("anda memasukkan angka 3");
//     break;
//   default:
//     alert("angka yang dimasukkan salah");
//     break;
// }

// let item = prompt(
//   "masukkan nama pesanan :\n (cnth: nasi, daging, susu, hamburger)"
// );

// switch (item) {
//   case "nasi":
//     alert("makanan/minuman sehat");
//     break;
//   case "daging":
//     alert("makanan/minumam sehat");
//     break;
//   case "susu":
//     alert("makanan/minuman sehat!");
//     break;
//   case "hamburger":
//     alert("makanan/minuman ");
//     break;
//   default:
//     alert("makanan/minumam tidak sehat!");
//     break;
// }

// let s = "";
// for (let i = 0; i < 5; i++) {
//   for (let j = 0; j < 10; j++) {
//     s += "*";
//   }
//   s += "\n";
//   // s += '*'
// }
// console.log(s);

let hasil1 = "";
for (let i = 0; i < 10; i++) {
  for (let j = 0; j <= i; j++) {
    hasil1 += "* ";
  }
  hasil1 += "\n";
}
console.log(hasil1);

let hasil2 = "";
for (let i = 0; i < 10; i++) {
  for (let j = 10; j > i; j--) {
    hasil2 += "* ";
  }
  hasil2 += "\n";
}
console.log(hasil2);

let hasil3 = "";
for (let i = 10; i > 0; i--) {
  for (let j = 0; j <= 10; j++) {
    if (j >= i) {
      hasil3 += "* ";
    } else {
      hasil3 += "  ";
    }
  }
  hasil3 += "\n";
}
console.log(hasil3);

let hasil4 = "";
for (let i = 10; i > 0; i--) {
  for (let j = 10; j > 0; j--) {
    if (j > i) {
      hasil4 += "  ";
    } else {
      hasil4 += "* ";
    }
  }
  hasil4 += "\n";
}
console.log(hasil4);
