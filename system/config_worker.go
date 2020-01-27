package system

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadConf() {
	f, e := os.Open("C:\\Users\\user\\go\\src\\thb\\config")
	if e != nil {
		panic(e)
	}

	conf := &Server{
		Handler: nil,
	}

	defer f.Close()

	sc := bufio.NewScanner(f)
	lineCount := 1
	for sc.Scan() {
		// Yorum satırları okunmaması için bu önlem alındı.
		if !strings.HasPrefix(sc.Text(), "##") && strings.ContainsRune(sc.Text(), '=') {
			if strings.Count(sc.Text(), "=") == 1 {
				line := strings.Split(sc.Text(), "=")

				key, value := strings.TrimSpace(line[0]), strings.TrimSpace(line[1])
				switch strings.ToLower(key) {
				case "port":
					conf.port = fmt.Sprintf(":%s", value) // :8080
				case "host":
					conf.host = value
				}
				lineCount += 1
			} else {
				tokIndex := strings.LastIndexByte(sc.Text(), '=')
				tok := string(sc.Text()[tokIndex])
				log.Panicf("Unexpected token: '%s' on conf file at: %d", tok, lineCount)
			}
		}
	}

	GetApplication().config = conf
}

//func GetConf() conf {
//	return GetApplication().config
//}
