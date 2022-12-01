import {Readable} from 'node:stream';
import * as readline from 'node:readline/promises';
import * as fs from 'node:fs/promises';


export function readableFromString(input){
    const stream = new Readable()
    stream.push(input);
    stream.push(null);
    return readline.createInterface(stream)
}

export async function run(testInput, expectedPart1, expectedPart2, inputFile, part1, part2){
    testInput = testInput.trim()
    const actualPart1 = await part1(readableFromString(testInput));
    const actualPart2 = await part2(readableFromString(testInput))
    console.log(`Expected Part 1 ${expectedPart1} = ${actualPart1}`)
    console.log(`Expected Part 2 ${expectedPart2} = ${actualPart2}`)

    const file1 = await fs.open(inputFile);
    const result1 = await part1(file1.readLines())
    const file2 = await fs.open(inputFile);
    const result2 = await part2(file2.readLines())
    console.log(`Result:\nPart 1: ${result1}\nPart 2: ${result2}`)
}