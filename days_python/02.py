def are_diffs_valid(diffs):
    return all(map(lambda x: x > 0 and x in range(1,4), diffs)) or all(map(lambda x: x < 0 and x in range(-3,0), diffs))

def is_report_valid_1(report):
    if len(report) == 1:
        return True
    diffs = [report[i] - report[i-1] for i in range(1, len(report))]
    if 0 in diffs:
        return False
    return are_diffs_valid(diffs)


def count_different_diffs(diffs):
    not_ascending = sum(x <= 0 or x not in range(1,4) for x in diffs)
    not_descending = sum(x >= 0 or x not in range(-3,0)for x in diffs)
    if not_ascending > not_descending:
        return not_descending, 'desc'
    return not_ascending, 'asc'

def remove_invalid_diff(diffs, direction):
    # FIXME: Well, there are just some edge cases that I can't seem to get right with this approach
    for i in range(len(diffs)):
        if (direction == 'asc' and (diffs[i] <= 0 or diffs[i] not in range(1,4))):
            new_diffs = diffs[:]
            if i < len(new_diffs) and i > 0:
                new_diffs[i-1] += diffs[i]
            new_diffs = new_diffs[:i] + new_diffs[i+1:]
            return new_diffs
        if (direction == 'desc' and (diffs[i] >= 0 or diffs[i] not in range(-3,0))):
            new_diffs = diffs[:]
            if i < len(new_diffs) and i > 0:
                new_diffs[i-1] += diffs[i]
            new_diffs = new_diffs[:i] + new_diffs[i+1:]
            return new_diffs
    return diffs

def is_report_valid_2(report):
    if len(report) == 1:
        return True
    diffs = [report[i] - report[i-1] for i in range(1, len(report))]
    diff_count, direction = count_different_diffs(diffs)
    if diff_count >= 1:
        # Check that Removing the element keeps valid difference
        new_diffs = remove_invalid_diff(diffs, direction)
        if are_diffs_valid(new_diffs):
            return True
        return False
    return False


if __name__ == '__main__':
    with open('../data/02.txt') as f:
        reports = [list(map(int, line.strip().split())) for line in f.readlines()]

    part_one, part_two = 0, 0
    for report in reports:
        if is_report_valid_1(report):
            part_one += 1
            part_two += 1
        elif is_report_valid_2(report):
            part_two += 1

    print(part_one, part_two)