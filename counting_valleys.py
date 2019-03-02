def countingValleys(n, s):
    sea_level = 0
    result = 0

    for i in range(n):
        if s[i] == 'U':
            sea_level += 1
        else:
            sea_level -= 1
        
        if sea_level == 0 and s[i]=='U':
            result += 1

    return result

n = 8
s = 'UDDDUDUU'
result = countingValleys(n, s)
print(result)