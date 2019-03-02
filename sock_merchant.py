# Find match pair of socks
def sockMerchant(n, ar):
    temp = {}
    result = 0

    for i in ar:
        if i in temp:
            temp[i] += 1
        else:
            temp[i] = 1
        
        if temp[i] == 2:
            result += 1
            temp[i] = 0
    
    return result


n = 9
ar = [10, 20, 20, 10, 10, 30, 50, 10, 20]
result = sockMerchant(n, ar)
print(result)