package banner

import (
	"fmt"

	"github.com/fatih/color"
)

func Print() {
	cyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgHiRed).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	fmt.Print(cyan(`
                                                                  
      * ***                                                     
    *  ****  *                                                  
   *  *  ****                                                   
  *  **   **                                                    
 *  ***          ***  ****                 ****    ***  ****    
**   **           **** **** *    ***      * ***  *  **** **** * 
**   **   ***      **   ****    * ***    *   ****    **   ****  
**   **  ****  *   **          *   ***  **    **     **         
**   ** *  ****    **         **    *** **    **     **         
**   ***    **     **         ********  **    **     **         
 **  **     *      **         *******   **    **     **         
  ** *      *      **         **        **    **     **         
   ***     *       ***        ****    * *******      ***        
    *******         ***        *******  ******        ***       
      ***                       *****   **                      
                                        **                      
                                        **                      
                                         **                     
                                                            
`))
	fmt.Println(green(bold("\nGrepr is a fast, flexible CLI tool for security researchers to filter important URLs by filetype, keywords, and patterns")))
	fmt.Println(red(bold("[+]Developer: Laviru Dilshan From Ovate Security[+]")))
	fmt.Println(yellow("GitHub: https://github.com/LaviruD"))
	fmt.Println(yellow("X: @laviru_dilshan"))
	fmt.Println(yellow("Linkdin: @laviru-dilshan\n"))
}
