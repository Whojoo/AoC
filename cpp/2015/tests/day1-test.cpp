#include <catch2/catch_test_macros.hpp>

unsigned int Factorial(unsigned int n)
{
	return n == 0 ? 1 : n * Factorial(n - 1);
}

TEST_CASE("Factorials are computed", "[factorial]") {
	REQUIRE(Factorial(1) == 1);
}