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

  var cap_lines = lines.join("\n");

  fs.writeFile("textfile_caps_nodejs.txt", cap_lines, function(err) {
    if (err) {
      return console.log(err);
    }
  });

  return cap_lines;
}

console.log(run());
