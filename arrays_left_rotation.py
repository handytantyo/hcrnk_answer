def rotLeft(a, d):
    rotate = d % len(a)
    result = [0] * len(a)
    for i in range(len(a)):
        checker = i+rotate
        if checker >= len(a):
            checker -= len(a)
        result[i] = a[checker]
    return result

n = 5
d = 4
a = [1,2,3,4,5]
result = rotLeft(a, d)
print(result)