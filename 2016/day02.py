keypad = [
    [1,2,3],
    [4,5,6],
    [7,8,9],
]
# starting location
x = 1
y = 1

def apply_instruction(line : str) -> None:
    global x, y
    for ins in line:
        newx, newy = x, y
        match ins:
            case "L":
                newx = newx - 1
            case "R":
                newx = newx + 1
            case "U":
                newy = newy - 1
            case "D":
                newy = newy + 1

        x = newx if newx >=0 and newx < len(keypad[0]) else x
        y = newy if newy >=0 and newy < len(keypad) else y
    print("eol", keypad[y][x])

keypad2= [
    ["", "", "1", "", ""],
    ["", "2", "3", "4", "",],
    ["5","6", "7", "8", "9"],
    ["", "A", "B", "C", ""],
    ["", "", "D", "", ""],
]
x2 = 0
y2 = 2

def apply_instruction_key2(line:str) -> None:
    global x2, y2
    for ins in line:
        newx, newy = x2, y2
        match ins:
            case "L":
                newx = newx - 1
            case "R":
                newx = newx + 1
            case "U":
                newy = newy - 1
            case "D":
                newy = newy + 1

        # clamping
        newx = max(0, min(len(keypad2[0])-1, newx))
        newy = max(0, min(len(keypad2)-1, newy))

        if keypad2[newy][newx]:
            x2, y2 = newx, newy

    print(keypad2[y2][x2])

if __name__ == "__main__":
    with open("day02.txt") as file:
        while line := file.readline():
            apply_instruction_key2(line)

