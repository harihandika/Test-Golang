console.log(test("Sample Case")); 
function test(e) { 
  var words = e 
    .toLowerCase() // for better compare and sort 
    .split("") // convert string to array 
    .filter((el) => el !== " "); // remove spaces 
 
  var vowel = words 
    .filter((el) => 
      "aiueo" // to find vowel letters 
        .split("") // to create vowel letter array 
        .includes(el) 
    ) 
    .sort( 
      (a, b) => a.localeCompare(b) // sort from a to z. you can replace a with b and b with a for reverse sorting 
    ) 
    .join(""); // to convert array to string 
 
  var consonant = words 
    .filter((el) => !"aiueo".split("").includes(el)) 
    .sort((a, b) => a.localeCompare(b)) 
    .join(""); 
 
  return { vowel, consonant }; 
}

// const prompt = require("prompt-sync")();

// let text = prompt("Input one line of words: ");
// text = text.toLowerCase().split(" ").join("");
// let vowel = "aiueo";
// let resultVowel = "";
// let resultConsonant = "";
// for (let i = 0; i < text.length; i++) {
//   if (vowel.includes(text[i])) {
//     resultVowel += text[i];
//   } else {
//     resultConsonant += text[i];
//   }
// }
// console.log("Vowel Characters :\n" + resultVowel);
// console.log("Consonant Characters :\n" + resultConsonant);