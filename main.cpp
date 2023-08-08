#include "common.hpp"
#include "lexer.hpp"
#include "parser.hpp"
#include "tokenizer.hpp"

int main(/*int argc, char* argv[]*/)
{
	clear_screen();

	std::cout << "Math Expression Evaluator" << "\n";
	std::cout << "Enter quit or exit or q to close the program" << "\n";
	std::string checkInput, expression;

	Parser parser;
	Lexer lexer;
	Tokenizer tokenizer;

	while (true)
	{
		std::cout << ">> ";
		std::getline(std::cin, expression);

		checkInput = toLowercase(expression);
		// check for exit commands
		if (checkInput == "exit" || checkInput == "quit" || checkInput == "q")
			break;

		auto tokens = tokenizer.tokenize(expression);

		auto lexedTokens = lexer.lex(tokens);

		double result = parser.parse(lexedTokens);
		if (!is_whole_number(result))
		{
			std::cout.precision(10);
			std::cout.setf(std::ios_base::fixed, std::ios_base::floatfield);
		}

		// Print the answer
		std::cout << result << '\n';
	}

	return 0;
}