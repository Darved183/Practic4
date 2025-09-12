package main

import (
	"Practic6/DataBase"
	"fmt"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var result = make(chan []string, 2)

func main() {

	Base := map[string][]string{

		"1": {"https://slazpvpfvhmqaomyivko.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InNsYXpwdnBmdmhtcWFvbXlpdmtvIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTc1ODY3MTUsImV4cCI6MjA3MzE2MjcxNX0.hGWqf-Meo6Bmmlgek9urcaEET79bJBZxePzFbRyamaY", "Франкфурт - Германия"},
		"2": {"https://wkbaokoeyshxukloffll.supabase.co", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6IndrYmFva29leXNoeHVrbG9mZmxsIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTc1OTA0OTUsImV4cCI6MjA3MzE2NjQ5NX0.Gp8sDvgweXzYnfAnC6Tf_zCkyrFVb4qa07i4Fi1HQTU", "Стокгольм - Швеция"},
	}

	fmt.Print("Введите ID интересуемой страны (например RU, EU, US): ")
	var Input string
	fmt.Scan(&Input)
	Input = strings.ToUpper(Input)
	for _, i := range Base {
		wg.Add(1)
		go func(i []string, Input string) {
			defer wg.Done()
			start := time.Now().UnixNano()
			client, err := DataBase.ConnectBase(i[0], i[1])
			if err != nil {
				fmt.Printf("Не удалось подключиться к базе данных: %v\n", err)
			}

			res := DataBase.SelectInBase(client, "Сountry", "id_country", Input)
			if len(res) > 0 {
				result <- res
			} else {
				fmt.Printf("Такой страны (%v) в базе нет!\n", Input)
			}
			elapsed := time.Duration(time.Now().UnixNano() - start)
			fmt.Printf("Время подключения к %v: %v\n", i[2], elapsed)
		}(i, Input)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	for i := range result {
		fmt.Println(i[0])
	}

}
