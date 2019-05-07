// bnech.js  Run inside of NodeJS

var fs = require("fs");

function capitalize(string) {
  return string.charAt(0).toUpperCase() + string.slice(1);
}

function processLine(line) {
  var words = line.split(" ");
  for (var i = 0; i < words.length; i++) {
    words[i] = capitalize(words[i]);
  }

  return words.join(" ");
}

function run() {
  var data = fs.readFileSync("textfile_1000kb.txt", "utf-8");
  var lines = data.split("\n");

  for (var i = 0; i < lines.length; i++) {
    lines[i] = processLine(lines[i]);
  }

  return lines.join("\n");
}

console.log(run());
