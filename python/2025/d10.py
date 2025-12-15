import re
import numpy as np
from scipy.optimize import linprog

lines = open("inputs/2025/d10.txt").read().splitlines()

total_sum = 0
for line in lines:
    # Parse tuples: (3,7) (0,1,3,4,7) ...
    vectors = [list(map(int, m.split(','))) for m in re.findall(r'\((.*?)\)', line)]

    # Parse curly braces: {15,28,22,56,34,25,25,28}
    result = list(map(int, re.search(r'\{(.*?)\}', line).group(1).split(',')))

    A=[]

    for v in vectors:
        row = [0 for _ in range(len(result))]
        for n in v:
            row[n] = 1
        A.append(row)

    matrix =np.column_stack(A)
    b = np.array(result, dtype=int)
    c = np.ones(len(A), dtype=int)

    integrality =[1 for _ in range(len(c))]
    result = linprog(c, A_eq=matrix, b_eq=b, bounds=(0, None), method='highs', integrality=integrality)

    if result.success:
        print(result.x)
        total_sum += sum(result.x)
    else:
        print("Optimization failed")




print(total_sum)
