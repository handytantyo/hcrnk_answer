def minimumSwaps(arr):
    panjang = len(arr)
    result = 0

    for i in range(panjang-1):
        while arr[i] != i+1:
            temp = arr[arr[i]-1]
            arr[arr[i]-1] = arr[i]
            arr[i] = temp
            result += 1
    return result

arr = [4, 3, 1, 2]
result = minimumSwaps(arr)
print(result)