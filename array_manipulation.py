def arrayManipulation(n, queries):
    arr = [0]*(n+1)
    for i in queries:
        a = i[0]
        b = i[1]
        k = i[2]
        arr[a-1] += k
        arr[b] -= k

    max_ = arr[0]
    for i in range(1, n):
        arr[i] += arr[i-1]
        max_ = max(arr[i],max_)

    return max_


n = 5
m = 3
queries = [[1, 2, 100],[2, 5, 100],[3, 4, 100]]
result = arrayManipulation(n , queries)
print(result)