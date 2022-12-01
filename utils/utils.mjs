import {Readable} from 'node:stream';
import * as readline from 'node:readline/promises';
import * as fs from 'node:fs/promises';
import assert from 'node:assert';


export function readableFromString(input){
    const stream = new Readable()
    stream.push(input);
    stream.push(null);
    return readline.createInterface(stream)
}

export async function run(inputFile, part1, part2){
    const file1 = await fs.open(inputFile);
    const result1 = await part1(file1.readLines())
    const file2 = await fs.open(inputFile);
    const result2 = await part2(file2.readLines())
    console.log(`Result:\nPart 1: ${result1}\nPart 2: ${result2}`)
}

export async function test(testInput, expected, part){
    testInput = testInput.trim()
    const actual = await part(readableFromString(testInput));
    assert.equal(actual, expected)
}