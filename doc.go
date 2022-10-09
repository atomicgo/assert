/*
Package assert provides a set of assertion functions.
Every assertion function returns a boolean. This package does not integrate into the testing package automatically.
If you want to use this package inside unit tests, you have to check the returning boolean value and call t.Fatal() if the assertion fails.

This library is on a very low level and does not provide any error messages. That way the assertions can also be used in production code.

If you want a full-featured assertion framework for unit tests, we recommend https://github.com/MarvinJWendt/testza
*/
package assert
