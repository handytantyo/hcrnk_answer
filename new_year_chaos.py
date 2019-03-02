def minimumBribes(q):
    result = 0
    swap = True
    
    while swap:
        swap = False
        for i in range(1, len(q)):
            if(q[i - 1] - i > 2):
                return 'Too chaotic'

            if(q[i - 1] > q[i]):
                temp = q[i]
                q[i] = q[i - 1]
                q[i - 1] = temp

                swap = True
                result = result + 1

    return result

a = [2, 1, 5, 3, 4]
b = [2, 5, 1, 3, 4]
print(minimumBribes(a))
print(minimumBribes(b))