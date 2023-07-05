# LinearRegressionComp

Linear Regression Comparison between R, Python, and Go

# Introduction

The Linear Regression Comp app is a Go application which compairs the results of a linear regression model built on the
Anscombe Quartet in Python, R, and Go.

## Features
1. Calculate the time and coefficients of each set of the Anscombe Quartet individually.
2. Calculate the time and coefficients of all of the Anscombe Quartet sets at the same time.
3. View average performance for all runs in current session.
4. (Linux Only!) View graph showing all runs.

## Notes

To time performance relative to each language, please note the following:

1. the anscombe quartet has been recreated in each languages native format.
2. the timeing experiment takes place at the platform level; Both Python's and R's functions are calculated in Python
   and R respectively, and then passed to the GO API. This ensures that the times are accurate and are not impacted by
   the API, time spent loading the platform, processing functions, etc. The timers ONLY calculate the exact time each function was
   ran.
3. By default, each "run" calculates the time to fit 500 models, and then the average time is calculated. This allows
   for the language to "warm up" as well as compensate for rounding issues when a model takes less time to train than
   the accuracy of the system clock.
4. By default, each "experiment" is performed 15 times, and the average time is shown. All of the indivual experiment
   results are saved.

# User Configurable Parameters

The following parameters can be set by the user by accessing the "Configuration Menu"

1. nRuns (Default: 500): Sets the number of times each model is fit.
2. nExperiments (Default: 15): Sets the number of times each experiment is run.
3. roundCoefficients (Default: 7): Number of decimal points to which the model coefficients are rounded.
4. roundTime (Default: 8): Number of decimal points to which the time in seconds is rounded.

# Requirements

Please ensure that the following packages have been installed within your current enviornment.

## 1. Python
* statsmodels.api
* json
* numpy

## 2. R
* jsonlite

## 3. Go
* github.com/guptarohit/asciigraph 
* github.com/inancgumus/screen
* github.com/montanaflynn/stats
* github.com/olekukonko/tablewriter


# Additional Experiment Notes

Note: the number of iterations run noticably decreased the time of the function. As a result of the study, the default
number of iterations has been set to 100 (from 10 initially).
