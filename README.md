# expreval: expression evaluator

A simple command-line tool for evaluating mathematical expressions.

This project provides a simple command-line tool for evaluating mathematical expressions. It includes a lexer and parser to analyze and calculate mathematical operations.

Development is active, and new features are being added. Feel free to contribute or suggest improvements!

## Table of Contents

- [Usage](#usage)
- [Support](#support)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

## Usage

Ensure you have Go installed. Clone the repository and run the following command:

```bash
go run main.go
```

Enter mathematical expressions when prompted, and the tool will evaluate and display the result. To exit the program, type `quit`, `exit` or `q`.

## Roadmap

- [] Support for more complex mathematical functions.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests. Before contributing, please follow these guidelines:

- Follow the [code of conduct](CODE_OF_CONDUCT.md).
- Keep a consistent coding style. To ensure your coding style remains the same, format your code with:
  ``` console
  $ go fmt path/to/source_code
  ```
- Use the stable version of Go.
- If you have to use an external package, please prefer lightweight ones (e.g., `fasthttp` over `net/http`).
- Prefer using the standard library over reinventing the wheel.

Please stick to Go's official style guidelines while submitting patches.
For support, please [open an issue](https://github.com/walker84837/expreval/issues).

## License

This project is licensed under the [BSD-3-Clause License](LICENSE.md).
