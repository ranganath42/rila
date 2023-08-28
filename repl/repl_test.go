package repl_test

import (
	"bufio"
	"bytes"
	"github.com/ranganath42/rila/repl"
	"strings"
	"testing"
)

func TestRepl(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
	x + y;
};
`
	expected := `>> LET        let
IDENT      five
=          =
INT        5
;          ;
>> LET        let
IDENT      ten
=          =
INT        10
;          ;
>> LET        let
IDENT      add
=          =
FUNCTION   fn
(          (
IDENT      x
,          ,
IDENT      y
)          )
{          {
>> IDENT      x
+          +
IDENT      y
;          ;
>> }          }
;          ;
>> `
	t.Run("TestRepl", func(t *testing.T) {
		in := strings.NewReader(input)
		var b bytes.Buffer
		out := bufio.NewWriter(&b)
		repl.Start(in, out)
		out.Flush()
		actual := b.String()
		if actual != expected {
			t.Errorf("Expected: %q, got: %q", expected, actual)
		}
	})
}
