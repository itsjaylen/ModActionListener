package handler

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mods []string

// Loading the users from the files
// TODO : implement the loading from a API when unbanned on twitch.
func LoadUsers() []string {
	file, err := os.Open("../../mods.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var mods []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		username := strings.TrimSpace(scanner.Text()) 
		if username != "" {                           
			mods = append(mods, username)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	//fmt.Println("Loaded mods:", strings.Join(mods, ", "))
	
	return mods

}

