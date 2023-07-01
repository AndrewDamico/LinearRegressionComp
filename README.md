# LinearRegressionComp
Linear Regression Comparison between R, Python, and Go

#TODO 

The Go, Python, and R scripts each contain their own anscombe sets. This ensures that the sets were run by the scripts in the native format. 
Testing times were also calculated by the platform which was being tested. This ensures that the results did not include time spent loading the platform, processing functions, etc. The timers ONLY calculate the exact time each function was ran.
Note for future testing. The timers calculated the time for each model to run once, and those times were then stored in a list, of which the average was taken. It would be beneficial to compare results against timing how long all iterations took and then dividing by the number of iterations.

Note: the number of iterations run noticably decreased the time of the function. As a result of the study, the default number of iterations has been set to 100 (from 10 initially).


