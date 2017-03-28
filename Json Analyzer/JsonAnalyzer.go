package main

import "io/ioutil"
import "fmt"
import "strings"
import "os"


func display(token, kind string) {
        fmt.Printf("%-15s %s\n", token, kind)
}

func splitAtLastDot(name string) (string, string) {
  dotLoc := strings.LastIndex(name, ".")
  if dotLoc == -1 { return name, "" }
  return name[:dotLoc], name[dotLoc + 1:]
}

func check_str_words(printedString string,charByte byte,i int ,fileCharacters string)(string,byte,int){  // to deal with strings that starts with "QUATATION"
			printedString += string(charByte)	// charByte is a single character 
													// printedString is accumlator of all the characters in whole string
				
				i++

				charByte = fileCharacters[i]
				
				for (charByte!=34){  
					printedString += string(charByte)
					charByte = fileCharacters[i]
	
					if(charByte==92){   // if / found, then next char must be "  , \ , / , b , f , n , r , t , u 
						i++ 
						charByte = fileCharacters[i]
						if(charByte==34 || charByte==98 || charByte==102 || charByte==114 || charByte==116 || charByte==92 || charByte==110 || charByte==47 || charByte==117) {
							printedString += string(charByte)
							i++
							charByte = fileCharacters[i]
						}
					} else{	
						i++ 							//incrementing i
						charByte = fileCharacters[i]
					}
				}
				printedString += string(charByte)  

				display(printedString, "STRING")

				
				return printedString,charByte,i
}

func characterChecker(fileCharacters string){

	var charByte byte
	for i:= 0; i<len(fileCharacters); i++ {
		charByte = fileCharacters[i]
		var printedString string
		switch {
			case charByte==116:
				// checking last character to be 'e'
			 	charByte=fileCharacters[i+3]  
			 	if(charByte==101){
			 		display("true", "TRUE")
			 	}
			case charByte==110:
				 // checking last character to be 'l'
			 	charByte=fileCharacters[i+3] 
			 	if(charByte==108){
			 		display("null", "NULL")
			 	}
			case charByte==102:
				// checking last character to be 'e'
			 	charByte=fileCharacters[i+4]  
			 	if(charByte==101){
			 		display("false", "FALSE")
			 	}
			case charByte==91:
				display("[","OPEN_BRACKET")
			case charByte==93:
				display("]","CLOSE_BRACKET")
			case charByte==123:
				display("{","OPEN_BRACE")
			case charByte==125:
				display("}","CLOSE_BRACE")
			case charByte==44:
				display(",","COMMA")
			case charByte==58: 
				display(":","COLON")
			case charByte==34: 	   // if " found
				printedString,charByte,i = check_str_words(printedString,charByte,i,fileCharacters) 				
			case (charByte>=48 && charByte<=57) || charByte==45 : // CHECKING ASCI OF 0-9 and negative 
			 	for ((charByte>=48 && charByte<=57) || charByte==43 || charByte==46 || charByte==69 || charByte==45 || charByte==101){
			 		printedString += string(charByte)
			 		i++
			 		charByte=fileCharacters[i]
			 	}
			 	i--
			 	display(printedString, "NUMBER")

			
		}
	}
}


func main(){


	nargs := len(os.Args)
  	if nargs != 2 {
    	fmt.Println("Error: exactly one filename argument must be provided")
    	panic("exited")
  	}

	inname := os.Args[1]
	infile, err := os.Open(inname)
	_, ext := splitAtLastDot(inname)   // to check the valid file type as this program is intend to run with json file as indicated in assignment
	
	if(ext!="json"){
		panic("Provided file is not a JSON file")
	}

	if err != nil {
	    fmt.Printf("Error: cannot open file \"%s\" (%s)", inname, err)
		panic("exited")
	}
	defer infile.Close()
	charContents, _ := ioutil.ReadFile(inname)
	fileCharacters := string(charContents)
	characterChecker(fileCharacters)	
}
