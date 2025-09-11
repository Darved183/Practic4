package main

import (
	"fmt"

	"github.com/supabase-community/postgrest-go"
)

func main() {

	URL1 := "https://slazpvpfvhmqaomyivko.supabase.co"
	URL2 := "https://wkbaokoeyshxukloffll.supabase.co"
	KEY1 := "yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InNsYXpwdnBmdmhtcWFvbXlpdmtvIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTc1ODY3MTUsImV4cCI6MjA3MzE2MjcxNX0.hGWqf-Meo6Bmmlgek9urcaEET79bJBZxePzFbRyamaY"
	KEY2 := "yJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6IndrYmFva29leXNoeHVrbG9mZmxsIiwicm9sZSI6ImFub24iLCJpYXQiOjE3NTc1OTA0OTUsImV4cCI6MjA3MzE2NjQ5NX0.Gp8sDvgweXzYnfAnC6Tf_zCkyrFVb4qa07i4Fi1HQTU"

	Base1 := postgrest.NewClient(URL1, KEY1, nil)
	Base2 := postgrest.NewClient(URL2, KEY2, nil)

	Search(Base1, "Россия") //Germany
	Search(Base2, "Россия") //Stocholm
}

func Search(client *postgrest.Client, productName string) {

	cursor := client.From("product").
		Select("*", "", false).
		Eq("name", productName)
	data, _, _ := cursor.Execute()
	fmt.Printf("Сырые данные: %s\n", string(data))

}
