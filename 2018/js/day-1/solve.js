var fs = require('fs')

class Runner {
  constructor () {
    this.input = fs
      .readFileSync('./input.txt')
      .toString()
      .replace(/ /g, '')
      .split(/\n/)
  }

  part1 () {
    return this.input.reduce((accumulator, str) => {
      const h = this.valueHash(str)

      h.operator === '-' ? (accumulator -= h.value) : (accumulator += h.value)

      return accumulator
    }, 0)
  }

  part2 () {
    const map = new Map()
    var accumulator = 0
    try {
      while (true) {
        this.input.forEach(str => {
          const h = this.valueHash(str)
          h.operator === '-' ? (accumulator -= h.value) : (accumulator += h.value) // can we use closures to extact this out?

          if (map.has(accumulator)) {
            throw new Error(accumulator)
          } else {
            map.set(accumulator, true)
          }
        })
      }
    } catch (err) {
      return err.message
    }
  }

  valueHash (str) {
    return { operator: str.substr(0, 1), value: parseInt(str.substr(1)) }
  }
}

const runner = new Runner()

console.log(runner.part1())
console.log(runner.part2())
