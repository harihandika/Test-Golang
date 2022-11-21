for(let x=1; x <= 10; x++) { 
 let temp = "" 
 let result = 1 
 console.log("ini yang x",x);
for(let y=1;y<=x;y++) {  
if(y < x) { 
result *= y 
temp = temp + y + " * " 
// console.log("ini yang y", temp); 
} else { 
result *= y 
temp = temp + y  
console.log("ini yang z", temp); 
} 
} 
console.log(temp + " = " + result) 
}