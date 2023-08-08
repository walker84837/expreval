#ifndef TOKENIZER_HPP_
#define TOKENIZER_HPP_

#include "common.hpp"

class Tokenizer
{
public:
	std::vector<std::string> tokenize(std::string expression)
	{
		std::vector<std::string> tokens;
		std::string token = "";
		std::string ops = "+-*/()^%";

		for (auto c : expression)
		{
			// Skip empty space
			if (c == ' ')
				continue;

			// handle operator
			if (ops.find(c) != std::string::npos)
			{
				if (!token.empty())
				{
					tokens.push_back(token);
				}

				token.clear();
				token += c;
				tokens.push_back(token);
				token.clear();
			}
			else
			{
				token += c;
			}
		}

		if (!token.empty())
		{
			tokens.push_back(token);
		}
		return tokens;
	}
};

#endif // TOKENIZER_HPP_