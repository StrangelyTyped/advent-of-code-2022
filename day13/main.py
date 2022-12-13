
import functools
import json

def cmp(a1, a2):
    if isinstance(a1, int) and isinstance(a2, int):
        return a1 - a2
    elif isinstance(a1, int):
        return cmp([a1], a2)
    elif isinstance(a2, int):
        return cmp(a1, [a2])
    else:
        # both lists
        i = 0
        while i < len(a1) and i < len(a2):
            x = cmp(a1[i], a2[i])
            if x != 0:
                return x
            i += 1
        return len(a1) - len(a2)


def part1(lines):
    idx = 1
    sum = 0
    i = 0
    while i < len(lines):
        line1 = lines[i].strip()
        line2 = lines[i+1].strip()
        i += 3
        
        if cmp(json.loads(line1), json.loads(line2)) <= 0:
            sum += idx
            print(str(idx) + " is sorted")
        else:
            print(str(idx) + " invalid")
        idx += 1
    print(sum)

def part2(lines):
    i = 0

    div1 = [[2]]
    div2 = [[6]]

    outLines = [div1,div2]

    while i < len(lines):
        line1 = lines[i].strip()
        line2 = lines[i+1].strip()
        i += 3
        outLines.append(json.loads(line1))
        outLines.append(json.loads(line2))
    outLines.sort(key=functools.cmp_to_key(cmp))
        
    print((outLines.index(div1) + 1) * (outLines.index(div2) + 1))
        
        

with open("../inputs/day13.txt","r") as f:
    lines = f.readlines()
    part1(lines)
with open("../inputs/day13.txt","r") as f:
    lines = f.readlines()
    part2(lines)

testInput = """
[1,1,3,1,1]
[1,1,5,1,1]

[[1],[2,3,4]]
[[1],4]

[9]
[[8,7,6]]

[[4,4],4,4]
[[4,4],4,4,4]

[7,7,7,7]
[7,7,7]

[]
[3]

[[[]]]
[[]]

[1,[2,[3,[4,[5,6,7]]]],8,9]
[1,[2,[3,[4,[5,6,0]]]],8,9]
""".strip().split("\n")
part1(testInput)
part2(testInput)