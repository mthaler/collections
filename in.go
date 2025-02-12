package collections

import (
	"bufio"
	"os"
)

/**
 *  The <code>In</code> data type provides methods for reading strings
 *  and numbers from standard input, file input, URLs, and sockets.
 *  <p>
 *  The Locale used is: language = English, country = US. This is consistent
 *  with the formatting conventions with Java floating-point literals,
 *  command-line arguments (via {@link Double#parseDouble(String)})
 *  and standard output.
 *  <p>
 *  For additional documentation, see
 *  <a href="https://introcs.cs.princeton.edu/31datatype">Section 3.1</a> of
 *  <i>Computer Science: An Interdisciplinary Approach</i>
 *  by Robert Sedgewick and Kevin Wayne.
 *  <p>
 *  Like {@link Scanner}, reading a token also consumes preceding Java
 *  whitespace, reading a full line consumes
 *  the following end-of-line delimiter, while reading a character consumes
 *  nothing extra.
 *  <p>
 *  Whitespace is defined in {@link Character#isWhitespace(char)}. Newlines
 *  consist of \n, \r, \r\n, and Unicode hex code points 0x2028, 0x2029, 0x0085;
 *  see <a href="http://www.docjar.com/html/api/java/util/Scanner.java.html">
 *  Scanner.java</a> (NB: Java 6u23 and earlier uses only \r, \r, \r\n).
 *
 *  @author David Pritchard
 *  @author Robert Sedgewick
 *  @author Kevin Wayne
 */
type In struct {
	scanner bufio.Scanner
}

/**
 * Initializes an input stream from a file.
 *
 * @param  file the file
 */

func NewFromFile(file *os.File) *In {
	return &In{scanner: *bufio.NewScanner(file)}
}

/**
 * Initializes an input stream from a filename
 *
 * @param  name the filename
 */
func NewFromFilename(name string) *In {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	return NewFromFile(file)
}
