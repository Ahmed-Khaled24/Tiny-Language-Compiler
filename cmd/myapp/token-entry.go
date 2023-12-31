package main

import (
	"fmt"

	"github.com/ahmedelsayed968/Compilers-Project/internal/Parser"
	"github.com/ahmedelsayed968/Compilers-Project/internal/Scanner"
)

func TestParser() *Parser.Node {
	test := Scanner.Fsm.Scan(`{ Sample program in TINY language - computes factorial}
	read x;   {input an integer }
	if  0 < x   then     {  don't compute if x <= 0 }
	   fact  := (1);
	   repeat
		  fact  := fact *  x;
		   x  := x  -  1
	   until x  =  0;
	   write  fact
	end`)

	var x []Scanner.Token
	for i := range test {
		// check if the token is a error then continue
		if test[i].TokenType == "ERROR" {
			continue
		}
		x = append(x, test[i])
	}
	// print tokens without errors (for now only)
	fmt.Println(x)

	parser := Parser.NewParser(x)
	parsedTokens := parser.Parse()

	// fmt.Println(parsedTokens.)
	// for x=parser.TreeEntry, x!=nil; x=x.Next  {

	// }
	return parsedTokens
}

func main() {
	// token := Scanner.CreateToken(Scanner.SPECIALSYMBOL, "=")
	// fmt.Println(*token)

	// Scanner.Fsm.Scan("read 3;    if  0 < x   then     fact  := 1;repeat fact  := fact *  x; x  := x  -  1until  x  =  0;write  fact  end")
	// test := Scanner.Fsm.Scan(`{ Sample program in TINY language – computes factorial}
	// read 3;   {input an integer }
	// if  0 < x   then     {  don’t compute if x <= 0 }
	//    fact  := 1;
	//    repeat
	// 	  fact  := fact *  x;
	// 	   x  := x  -  1
	//    until {rerer} x  =  0;
	//    write  fact
	// end`)

	// fmt.Println(test)

	TestParser()

}
