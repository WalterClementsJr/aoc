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
        # print("step", x,y,keypad[y][x])

    print("eol", keypad[y][x])

def apply_instruction_key2(line:str) -> None:
    pass
if __name__ == "__main__":
    inputs = [
    "ULL",
    "RRDDD",
    "LURDL",
    "UUUUD",
    ]
    with open("day02.txt") as file:
        while line := file.readline():
        # for line in inputs:
            # print(line)
            apply_instruction(line)



