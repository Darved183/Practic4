package DataBase

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

type Country struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Id_country string `json:"id_country"`
}

func ConnectBase(URL string, Key string) (*supabase.Client, error) {
	client, err := supabase.NewClient(URL, Key, nil)
	if err != nil {
		return nil, fmt.Errorf("не удалось создать клиент Supabase: %w", err)
	}
	return client, nil
}

func SelectInBase(client *supabase.Client, table string, column string, value string) []string {
	var Countries []Country
	_, err := client.From("Country").Select("*", "", false).Eq(column, value).ExecuteTo(&Countries)
	if err != nil {
		fmt.Printf("Ошибка при запросе данных из таблицы %s: %v", table, err)
		return []string{}
	}
	if len(Countries) == 0 {
		return []string{}
	}

	return []string{Countries[0].Name, Countries[0].Id_country}

}
