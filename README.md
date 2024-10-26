# Guess number
Guess number in a range(lower, upper) with given propotions

# How to use
1. Clone the project
2. Modify the input.json
3. Run the code and get the possible results
```
go run .
```

# Inspiration
While I search the job from [104](https://www.104.com.tw/), for each job, it only shows the range of applicants(<5, 5-10, 10-30, 30+).
![alt text](image.png)

## How to get the real number of applicants?
![alt text](image-1.png)
![alt text](image-2.png)

The clues are:
**For each proportion, there must have some (number / total) == proportion**

Ex.

1 / 9 &#8776; 11%

6 / 9 &#8776; 67%

So, we can use these clues to esitmate the real number of applications