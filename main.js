function countCharacters(inputStr) {
  const counts = {};
  
  for (let i = 0; i < inputStr.length; i++) {
    const char = inputStr.charAt(i);
    counts[char] = (counts[char] || 0) + 1;
  }
  
  return counts;
}

function printCharacterCounts(inputStr) {
  const counts = countCharacters(inputStr);
  let outputStr = "";
  
  for (const char in counts) {
    if (counts.hasOwnProperty(char)) {
      const count = counts[char];
      if (count === 1) {
        outputStr += char;
      } else {
        outputStr +=  count + char;
      }
    }
  }
  
  return outputStr;
}
function removeSpaces(inputStr) {
  return inputStr.replace(/\s/g, "");
}
function changetolowercase(input){
    return input.toLowerCase();
}
function main(input){
    return printCharacterCounts(removeSpaces(changetolowercase(input)));
}
const input = "dani Maulana";
console.log("text yang anda masukkan adalah: "+ input);
const output =  main(input);
console.log("Hasil setelah di kelola: "+ output);

const input2 = "Mas Syahdan Filsafan";
console.log("text yang anda masukkan adalah: "+ input2);
const output2 =  main(input2);
console.log("Hasil setelah di kelola: "+ output2);
