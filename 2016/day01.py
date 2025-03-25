from enum import Enum
from re import error


INPUT = "R4, R4, L1, R3, L5, R2, R5, R1, L4, R3, L5, R2, L3, L4, L3, R1, R5, R1, L3, L1, R3, L1, R2, R2, L2, R5, L3, L4, R4, R4, R2, L4, L1, R5, L1, L4, R4, L1, R1, L2, R5, L2, L3, R2, R1, L194, R2, L4, R49, R1, R3, L5, L4, L1, R4, R2, R1, L5, R3, L5, L4, R4, R4, L2, L3, R78, L5, R4, R191, R4, R3, R1, L2, R1, R3, L1, R3, R4, R2, L2, R1, R4, L5, R2, L2, L4, L2, R1, R2, L3, R5, R2, L3, L3, R3, L1, L1, R5, L4, L4, L2, R5, R1, R4, L3, L5, L4, R5, L4, R5, R4, L3, L2, L5, R4, R3, L3, R1, L5, R5, R1, L3, R2, L5, R5, L3, R1, R4, L5, R4, R2, R3, L4, L5, R3, R4, L5, L5, R4, L4, L4, R1, R5, R3, L1, L4, L3, L4, R1, L5, L1, R2, R2, R4, R4, L5, R4, R1, L1, L1, L3, L5, L2, R4, L3, L5, L4, L1, R3"

class CardinalDirection(Enum):
    N = 0
    E = 1
    S = 2
    W = 3


def absolute(n):
    return n if n > 0 else n*-1

def applyCommand(current_direction, current_location, command, steps, is_turning=False):
    """
    calculate new location and new facing direction
    @return array, CardinalDirection
    """
    x, y = current_location
    new_dir = current_direction

    if (is_turning):
        if command == "R":
            new_dir = current_direction.value + 1
        elif command == "L":
            new_dir = current_direction.value - 1
        else:
            raise error("Command not valid, R or L")

        new_dir = CardinalDirection(new_dir%len(CardinalDirection))

    # Y axis
    if new_dir.value % 2 == 0:
        grid_step = 1 if new_dir == CardinalDirection.N else -1
        return [x, y + steps * grid_step], new_dir
    else:
        # X axix
        grid_step = 1 if new_dir == CardinalDirection.E else -1
        return [x + steps * grid_step, y], new_dir

locationMap = {}
currentDirection = CardinalDirection.N
currentLocation = [0,0]

for command in INPUT.split(", "):
    dir = command[0]
    step = command[1:]
    step_count = int(step)
    end = False
    print("command", command)

    for i in range(1, step_count+1):
        print("\t", i, "current location", currentLocation, "curretnDirection", currentDirection)
        is_turning = True if i==1 else False

        currentLocation, currentDirection = applyCommand(currentDirection, currentLocation, dir, 1, is_turning)

        point = f"{currentLocation[0]}:{currentLocation[1]}"
        print("\t", i, "current location", currentLocation, "curretnDirection", currentDirection)

        if locationMap.get(point):
            print("found visited", currentLocation)
            print(absolute(currentLocation[0] + absolute(currentLocation[1])))
            end = True
            break

        locationMap.update({point: 1})
    if end:
        break


print("new location and new direction:", currentLocation, currentDirection)
print(absolute(currentLocation[0]) + absolute(currentLocation[1]))

