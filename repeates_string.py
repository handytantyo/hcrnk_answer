# extra
def maximumrepeatedString(s, n):
    temp = {}
    for i in s:
        if i not in temp:
            temp[i] = 1
        else:
            temp[i] += 1
    print(temp)

    # find multiplier and mod
    multiplier = n // len(s)
    mod = n % len(s)

    maximum_number = 0
    # insert multiplier into set
    for i in temp:
        temp[i] *= multiplier
        maximum_number = max(maximum_number, temp[i])
    
    # insert mod
    if mod > 0:
        for i in range(mod):
            temp[s[i]] += 1
            maximum_number = max(maximum_number, temp[s[i]])
    
    return maximum_number

def repeatedString(s, n):
    if 'a' not in s:
        return 0
    else:
        temp = {}
        for i in s:
            if i not in temp:
                temp[i] = 1
            else:
                temp[i] += 1

        # find multiplier and mod
        multiplier = n // len(s)
        mod = n % len(s)

        # insert multiplier into set
        for i in temp:
            temp[i] *= multiplier
        
        # insert mod
        if mod > 0:
            for i in range(mod):
                temp[s[i]] += 1
        
    return temp['a']


s = 'gfcaaaecbg'
n = 547602
result = repeatedString(s, n)
print(result)
