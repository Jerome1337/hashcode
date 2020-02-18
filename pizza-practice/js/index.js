const fs = require('fs'),
      path = require('path'),
      readline = require('readline');

const inputFolder = '../input/';

let maxSlice = 0;

async function start() {
  for (let file of fs.readdirSync(inputFolder)) {
    const pizzaSliceNum = await readFile(file);

    const pizzaType = await solve(pizzaSliceNum);

    await writeFile(file, pizzaType.reverse());
  };

  console.log('Finished input files processing.');
}

function readFile(file) {
  const rd = readline.createInterface({
    input: fs.createReadStream(path.join(inputFolder, file))
  });

  let pizzaSliceNum = [],
      isFirstLine = true;

  return new Promise(resolve => {
    rd.on('line', (line => {
      const array = line.split(" ");

      if (isFirstLine) {
        maxSlice = array[0];

        isFirstLine = !isFirstLine;
      }

      pizzaSliceNum = array;
    }));

    rd.on('close', (() => {
      resolve(pizzaSliceNum);
    }));
  })
}

function solve(pizzaSliceNum) {
  return new Promise(resolve => {
    let slicesSum = 0,
    pizzaType = [],
    sliceLen = pizzaSliceNum.length - 1;

    for (let pizza of pizzaSliceNum) {
      let tmpSliceSum = 0,
          tmpPizzaType = [];

      for (let i = sliceLen; i >= 0; i--) {
        let tmpSum = parseFloat(tmpSliceSum) + parseFloat(pizzaSliceNum[i]);

        if (tmpSum > maxSlice) {
          continue;
        }

        tmpSliceSum = tmpSum;
        tmpPizzaType.push(i);
      }

      if (tmpSliceSum > slicesSum) {
        slicesSum = tmpSliceSum;
        pizzaType = tmpPizzaType;
      }

      if (slicesSum == maxSlice) {
        break;
      }

      sliceLen--
    }

    resolve(pizzaType);
  });
}

function writeFile(inputFile, pizzaType) {
  return new Promise((resolve, reject) => {
    fs.writeFile(`./submission/${inputFile.replace("in", "sub")}`, `${pizzaType.length}\n${pizzaType.join(' ')}`, (err => {
      if (err) {
          reject("Error writting string", err);
      }

      resolve("File successfully written");
    }));
  });
}

start();
