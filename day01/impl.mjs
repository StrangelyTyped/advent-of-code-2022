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

export async function part1(readStream){
    return solve(readStream, 1);
}

export async function part2(readStream){
    return solve(readStream, 3);
}
