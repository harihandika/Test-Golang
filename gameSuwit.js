let tanya = true;
while (tanya) {

  // pilihan player
  let p = prompt("pilih : gajah, semut, orang");

  //pilihan computer
  // bilangan random
  let comp = Math.random();

  if (comp < 0.33) {
    comp = "gajah";
  } else if (comp >= 0.33 && comp <= 0.66) {
    comp = "semut";
  } else {
    comp = "orang";
  }

  let hasil = "";

  if (p == comp) {
    hasil = "SERI";
  } else if (p == "gajah") {
    //   if (comp == "orang") {
    //     hasil = "MENANG";
    //   } else {
    //     hasil = "KALAH";
    //   }
    hasil = (comp == "orang" )? "MENANG" : "KALAH";
  } else if (p == "orang") {
    //   if (comp == "gajah") {
    //     hasil = "KALAH";
    //   } else {
    //     hasil = "MENANG";
    //   }
    hasil = (comp == "semut") ? "MENANG" : "KALAH";
  } else if (p == "semut") {
    //   if (comp == "gajah") {
    //     hasil = "MENANG";
    //   } else {
    //     hasil = "KALAH";
    //   }
    hasil = (comp == "orang") ? "KALAH" : "MENANG";
  } else {
    hasil = "Pilihan salah";
  }
  alert('Kamu memilih : ' + p + ' dan Komputer memilih : ' + comp + '\nMaka hasilnya : Kamu ' + hasil );
  console.log(hasil);
  // menentukan rules

 tanya = confirm('lagi?');
}

alert('terimakasih sudah bermain')