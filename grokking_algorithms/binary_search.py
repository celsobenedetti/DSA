import math


def binary_search(input, target):
    low = 0
    high = len(input) - 1

    while low <= high:
        mid = math.floor((high + low) / 2)

        if input[mid] == target:
            return mid

        if input[mid] > target:
            high = mid - 1
        else:
            low = mid + 1

    return -1


my_list = [1, 3, 5, 7, 9]

print(binary_search(my_list, 5))
print(binary_search(my_list, -1))
