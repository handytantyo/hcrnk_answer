# Complete the hourglassSum function below.
def hourglassSum(arr):
    maximum_value = 0
    result = 0
    for i in range(len(arr)-2):
        for j in range(len(arr)-2):
            result = arr[i][j] + arr[i][j+1] + arr[i][j+2] + arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
            if i == 0 and j == 0:
                maximum_value = result
            elif result > maximum_value:
                maximum_value = result
            result = 0
    return maximum_value

arr = [[1, 1, 1, 0, 0, 0],
        [0, 1, 0, 0, 0, 0],
        [1, 1, 1, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 0]]

result = hourglassSum(arr)
print(result)