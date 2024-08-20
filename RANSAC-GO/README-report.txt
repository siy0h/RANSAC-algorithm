Hello!

This code implements the fan-in method, which is used to combine data from multiple channels into one output channel.
However, I was not able to figure out how to use multithreading in the implementation, so it currently runs on a single thread.
As a result, it may be slower than it could be with multiple threads.
Additionally, there are a number of hard-coded values in the code, such as the epsilon value, confidence interval, and filename. 
These values could be replaced with command line arguments to make the code more flexible.
Overall, this code may be useful for combining data from multiple channels,
but it may not be the most efficient implementation due to the lack of multithreading and hard-coded values.