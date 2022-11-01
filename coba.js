for(let x=1; x <= 10; x++) { 
 let temp = "" 
let result = 1 
for(let y=1;y<=x;y++) {  
if(y < x) { 
result *= y 
temp = temp + y + " * "  
} else { 
result *= y 
temp = temp + y  
} 
} 
console.log(temp + " = " + result) 
}