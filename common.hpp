#ifndef COMMON_HPP_
#define COMMON_HPP_

#include <string>
#include <vector>
#include <iostream>
#include <sstream>
#include <cmath>
#include <cstdlib>
#include <map>
#include <stack>
#include <stdexcept>
#include <ios>

/**
 * @brief Checks if a given number is whole or not.
 * 
 *
 * @param x Input variable (double).
 * @return 
 */
inline bool is_whole_number(double x)
{
	return x == std::floor(x);
}

inline void clear_screen()
{
	try {
#ifdef _WIN32
		std::system("CLS");
#else
		std::system("clear");
#endif
	}

	catch (std::exception& e) {
		std::cerr << "Failed to clear screen. Error: " << e.what() << '\n';
	}

}

inline std::string toLowercase(const std::string &input)
{
	std::string result = input;
	for (char &c : result)
	{
		c = std::tolower(c);
	}
	return result;
}

#endif // COMMON_HPP_