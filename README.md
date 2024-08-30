## How to Use the Go Program for Basic Arithmetic Operations

**Purpose:** This Go program is designed to perform basic arithmetic operations (addition, subtraction, multiplication, and division) on two numbers based on user input.

**Steps:**

1. **Compile the Program:**

   - Save the code in a `.go` file (e.g., `calculator.go`).
   - Open your terminal or command prompt and navigate to the directory where the file is saved.
   - Use the `go build` command to compile the program:
     ```bash
     go build calculator.go
     ```
   - This will create an executable file (e.g., `calculator` on Windows or `calculator` on Linux/macOS).

2. **Run the Program:**

   - Execute the generated executable file:
     ```bash
     ./calculator
     ```

3. **Input Values:**

   - The program will prompt you to enter:
     - The first number
     - The second number
     - The operator (+, -, \*, /)

4. **Output:**
   - The program will display the result of the calculation based on the entered values and operator.
   - If you attempt to divide by zero, an error message will be displayed.

**Example Usage:**

```
Enter the first number: 10
Enter the second number: 5
Enter the operator (+, -, *, /): +
15
```

**Explanation:**

- The program declares three variables: `num1`, `num2` (both of type `float64` to handle decimal numbers) and `operator` (of type `string`).
- It prompts the user to input the first number, second number, and operator.
- Using a `switch` statement, it checks the operator and performs the corresponding calculation.
- The result is then printed to the console.

**Additional Notes:**

- You can modify the program to handle more complex calculations or provide additional features as needed.
- Consider using error handling mechanisms to gracefully handle invalid input or unexpected conditions.
- For more advanced arithmetic operations or mathematical functions, you might explore using Go's built-in `math` package.
