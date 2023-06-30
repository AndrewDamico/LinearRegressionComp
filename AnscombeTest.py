#!/usr/bin/env python

#import pandas as pd
import numpy as np
import statsmodels.api as sm
#import matplotlib.pyplot as plt
import time
import sys
import json

anscombe = {
    "One":{
        'x':(10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5),
        'y':(8.04, 6.95,  7.58, 8.81, 8.33, 9.96, 7.24, 4.26,10.84, 4.82, 5.68)
    },
    "Two":{
        'x': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        'y': [9.14, 8.14,  8.74, 8.77, 9.26, 8.1, 6.13, 3.1,  9.13, 7.26, 4.74]
    },    
    "Three":{
        'x': [10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5],
        'y': [7.46, 6.77, 12.74, 7.11, 7.81, 8.84, 6.08, 5.39, 8.15, 6.42, 5.73]
    },    
    "Four":{
        'x': [8, 8, 8, 8, 8, 8, 8, 19, 8, 8, 8],
        'y': [6.58, 5.76, 7.71, 8.84, 8.47, 7.04, 5.25, 12.5, 5.56, 7.91, 6.89]
    },
}

def linearregression(x,y, params = True):
    ''' 
    returns linear regression model
    if params = True, returns coefficients only
    '''
    matrix = sm.add_constant(x)
    model = sm.OLS(y, matrix)
    res = model.fit()
    if params:
        res = res.params
        
    return res

def test(dataset):
    x = dataset['x']
    y = dataset['y']
    res = linearregression(x,y)
    return (res)

def timer(function, testset, n = 10):
    '''
    returns average function execution time for n runs
    '''
    times = []

    for i in range(n):
        start = time.process_time()
        res = function(testset)
        end = time.process_time() - start
        times.append(end)

    response = {
        "Coefficients": res.tolist(),
        "Time": np.mean(times)
        }
    json_response = json.dumps(response)

    return (json_response)


if __name__ == "__main__":
    set_n = sys.argv[1]
    print (timer(test, anscombe[set_n]))
    