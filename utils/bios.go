package utils

import (
	"bufio"
	"os"
	"regexp"

	"github.com/djlechuck/recalbox-manager/structs"
)

// GetBiosList returns list of all BIOS file with their valid MD5.
func GetBiosList(md5File string) []structs.BiosFile {
	re := regexp.MustCompile("^([a-f0-9]{32})[ ]+(.*)$")
	file, err := os.Open(md5File)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	list := []structs.BiosFile{}
	tmp := map[string][]string{}

	for scanner.Scan() {
		m := re.FindStringSubmatch(scanner.Text())

		if 3 != len(m) {
			continue
		}

		// m[1] -> MD5 ; m[2] -> BIOS name
		tmp[m[2]] = append(tmp[m[2]], m[1])
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// k -> BIOS name ; v -> MD5 list
	for k, v := range tmp {
		list = append(list, structs.BiosFile{Name: k, Md5: v})
	}

	return list
}
