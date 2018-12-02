var fs = require('fs')

class Runner {
  constructor() {
    this.input = fs
      .readFileSync("./input.txt")
      .toString()
      .replace(/ /g, "")
      .split(/\n/);
  }

  part1() {
    const checkObj = this.input.reduce(
      (accumulator, value) => {
        const enumuerator = this.enumerator(value);

        if (Object.values(enumuerator).includes(2)) {
          accumulator.two++;
        }

        if (Object.values(enumuerator).includes(3)) {
          accumulator.three++;
        }

        return accumulator;
      },
      { two: 0, three: 0 }
    );

    return checkObj.two * checkObj.three;
  }

  enumerator(string) {
    return string.split("").reduce((acc, val) => {
      val in acc ? (acc[val] += 1) : (acc[val] = 1);
      return acc;
    }, {});
  }

  part2() {
    const stringsToArray = this.input.map(val => val.split(""));

    return stringsToArray.reduce(
      (acc, checking) => {
        stringsToArray.forEach(potentialMatch => {
          const matchedChars = this.matchedStrings(checking, potentialMatch);

          if (!acc.includes(matchedChars)) {
            acc.push(matchedChars);
          }
        });
        return acc;
      },
      []
    ).join("");
  }

  matchedStrings(checkingStringAsArray, potentialStringAsArray) {
    const obj = checkingStringAsArray.reduce(
      (acc, char, index) => {
        if (char !== potentialStringAsArray[index]) {
          acc.matches++;
        } else {
          acc.string += char;
        }
        return acc;
      },
      { matches: 0, string: "" }
    );

    return obj.matches == 1 ? obj.string : ""
  }
}

const runner = new Runner()

console.log(runner.part1())
console.log(runner.part2())
