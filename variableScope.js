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

function faktorial (n) {
    if (n === 0 ) return 1;
    console.log(n)
    return n * faktorial(n-1)
}

console.log(faktorial(5));