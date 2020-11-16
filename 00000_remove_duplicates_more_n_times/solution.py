#!/usr/bin/python3

def answer(data, n):
    cnt = {}
    for i in range(len(data)):
        cnt[data[i]] = cnt.get(data[i],0)+1
    remove = []
    for k,v in cnt.items():
        if(v>n):
            for i in range(v):
                data.remove(k)
    return data

data = [5, 10, 15, 10, 7] # replace this with any other list
n = 1 # replace this with any other integer
print(answer(data, n))

data = [10,2,2,3,8,4,5,11,5,8] # replace this with any other list
n = 1 # replace this with any other integer
print(answer(data, n))