/*
 splitPrint.go V1.R0.M0

 This file is part of the splitPrint distribution.
 Copyright (c) 2023 James Salvino.
 
 This program is free software: you can redistribute it and/or modify  
 it under the terms of the GNU General Public License as published by  
 the Free Software Foundation, version 3.

 This program is distributed in the hope that it will be useful, but 
 WITHOUT ANY WARRANTY; without even the implied warranty of 
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU 
 General Public License for more details.

 You should have received a copy of the GNU General Public License 
 along with this program. If not, see <http://www.gnu.org/licenses/>.

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"runtime"
	"strings"
)

func main() {
	var tempName string
	var jobName string
	var jobNumber string
	var jobTime string
	var jobNameNumberTime string
	var line string
	var t string
	var pth string
	var endCtr int = 0
	var needToOpen bool = true
	var fo *bufio.Writer
	// var tfile *os.File

	platform := runtime.GOOS
	if platform != "windows" {
		tempName = "prt/JOBX.TXT"
		pth = "prt/"
	} else {
		tempName = "prt\\JOBX.TXT"
		pth = "prt\\"
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line = scanner.Text()
		
		if needToOpen {
			tfile, err := os.Create(tempName)
			if err != nil {
				fmt.Println("splitPrint: Create temp error: ", err)
			}
			defer tfile.Close()
			
			fo = bufio.NewWriter(tfile)
			needToOpen = false
		}

		_, err := fo.WriteString(line + "\n")
		if err != nil {
			fmt.Println("splitPrint: Write temp error: ", err)
		}

		if strings.HasPrefix(line, "****A   END") {
			endCtr += 1

			if endCtr == 4 {
				jobName = strings.TrimRight(line[24:32], " ")
				jobNumber = strings.ReplaceAll(line[14:22], " ", "0")
				t = strings.ReplaceAll(line[67:75], ".", "_")
				jobTime = strings.ReplaceAll(t, " ", "0")
				jobNameNumberTime = pth + jobName + "-" + jobNumber + "-" + jobTime + ".txt"

				fo.Flush()
				
				/*
				Note:
				Unable to use os.Rename() because the rename would fail under Windows with the following error: 
				"The process cannot access the file because it is being used by another process."
				os.Rename() worked fine on the Unix-like platforms
				To get around the problem - needed to copy the temp file to a newly created file 
				with the desired name
				*/
				
				source, err := os.Open(tempName)
				if err != nil {
					fmt.Println("splitPrint: Open src error: ", err)
				}
				defer source.Close()
		
				destination, err := os.Create(jobNameNumberTime)
				if err != nil {
					fmt.Println("splitPrint: Create dest error: ", err)
				}
				defer destination.Close()
		
				_, err = io.Copy(destination, source)
				if err != nil {
					fmt.Println("splitPrint: Copy error: ", err)
				}
						
				needToOpen = true
				endCtr = 0
			}
		}

	}

}
