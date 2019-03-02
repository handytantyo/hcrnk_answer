def jumpingOnClouds(c):
    angka = len(c)
    index = 0
    result = 1

    while index < angka:
        try:
            if c[index+1] == 0 and (index+1) == (angka-1):
                result+=1
                index += 1
            elif c[index+1] == 0 and c[index+2] == 0:
                result += 1
                index += 2
            elif c[index+1] == 0 and c[index+2] == 1:
                result += 1
                index += 1
            elif c[index+1] == 1 and c[index+2] == 0:
                result+= 1
                index += 2
        except:
            return result-1
    
    return result-1
        

n = 6
c = [0, 0, 0, 1, 0, 0]
result = jumpingOnClouds(c)
print(result)