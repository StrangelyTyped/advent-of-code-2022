import * as utils from '../utils/utils.mjs';

async function* parseElves(readStream){
    let elf = 0;
    for await (const line of readStream) {
        if(!line.trim().length){
            yield elf;
            elf = 0;
        } else {
            elf += parseInt(line)
        }
    }
    yield elf;
}

function sortedInsert(arr, n){
  let i = 0;
  while(arr[i] < n){
    i++;
  }
  arr.splice(i, 0, n)
  return arr;
}

async function solve(readStream, n){
    let vals = [];

    for await (const elf of parseElves(readStream)){
        if (vals.length < n){
            vals = sortedInsert(vals, elf)
        } else if (elf > vals[0]){
            vals = sortedInsert(vals.slice(1), elf)
        }
        
    }
    return vals.reduce((accum, i) => accum + i, 0)
}


async function part1(readStream){
    return solve(readStream, 1);
}

async function part2(readStream){
    return solve(readStream, 3);
}



const testInput = `
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

utils.run(testInput, 24000, 45000, "../inputs/day01.txt", part1, part2)
