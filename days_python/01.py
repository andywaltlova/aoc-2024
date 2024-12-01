def part1(left, right):
    return sum([abs(left[i] - right[i]) for i in range(len(left))])

def part2(left, right):
    right_occurences = {}
    for num in right:
        right_occurences[num] = right_occurences.get(num, 0) + 1

    return sum([num * right_occurences.get(num, 0) for num in left])

if __name__ == '__main__':
    with open('../data/01.txt') as f:
        lines = [list(map(int, line.split())) for line in f.readlines()]

    left = sorted(map(lambda x: x[0], lines))
    right = sorted(map(lambda x: x[1], lines))

    print(part1(left, right))
    print(part2(left, right))
