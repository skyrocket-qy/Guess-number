# Guess Number of Applicants

A Go program to deduce the possible number of applicants based on statistical data.

## The Story

Have you ever seen statistics on a job posting site and wondered how many people actually applied?

For example, you might see a range of applicants:

![Number of Applicants](image.png)

And you might also see the distribution of their educational backgrounds:

![Education Distribution](image-1.png)

This project helps you take these percentages and the applicant range, and determine the possible number of total applicants.

## How to Use

1.  **Edit the `input.json` file:**

    ```json
    {
      "lowerBound": 15,
      "upperBound": 25,
      "proportions": [
        [22, 67, 11],
        [45, 55]
      ]
    }
    ```

    *   `lowerBound`: The lower bound of the possible number of applicants.
    *   `upperBound`: The upper bound of the possible number of applicants.
    *   `proportions`: An array of proportion sets. Each proportion set is an array of percentages.

2.  **Run the program:**

    ```bash
    go run main.go
    ```

3.  **View the output:**

    The program will print the possible number(s) of applicants. For the example `input.json` above, the output will be:

    ```
    Possible numbers: [18]
    ```

## The Logic

The program works by iterating through each integer number in the range from `lowerBound` to `upperBound`. For each number, it checks if it satisfies all the given proportion sets.

A number `n` satisfies a proportion set if, for each percentage `p` in the set, the value `(p/100) * n` is very close to a whole number. This is because the number of people in any category must be an integer.

### Limitation

Please note that this program can only handle **exclusive proportions**, where the sum of the percentages in each set is approximately 100%. It cannot, for example, process data where one person can belong to multiple categories (like a skill distribution where percentages might sum to more than 100%). This is due to a validation check in the code that ensures the integrity of the proportions.
