CXX = g++
CXXFLAGS = -std=c++17 -O3 -Wall -Wextra -Wpedantic

SRC = main.cpp lexer.hpp parser.hpp tokenizer.hpp
TARGET = math_expression_evaluator

.PHONY: all clean

all: $(TARGET)

$(TARGET): $(SRC)
	$(CXX) $(CXXFLAGS) $^ -o $@

clean:
	@rm -f $(TARGET)
