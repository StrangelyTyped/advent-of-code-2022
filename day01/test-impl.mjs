import test from 'node:test';
import * as utils from '../utils/utils.mjs';

import { part1, part2 } from './impl.mjs';

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

test('provided sample part 1', async (t) => {
    await utils.test(testInput, 24000, part1)
});

test('provided sample part 2', async (t) => {
    await utils.test(testInput, 45000, part2)
});
  