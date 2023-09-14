package main

import (
	"IsAdmin/cache"
	"fmt"
	"time"
)

func main() {
	// Тестируем работу кэша
	// Создаем контейнер с временем жизни по-умолчанию равным 5 минут и удалением просроченного кеша каждые 10 минут
	Mycache := cache.New(5*time.Minute, 10*time.Minute)

	// Установить кеш с ключем 365557(uint64) и временем жизни 5 минут
	usr := cache.NewUser(365557)
	Mycache.Set(usr.GetUser(), true, 5*time.Minute)

	// Получить кеш с ключем 365557
	i, _ := Mycache.Get(usr.GetUser())

	fmt.Println(i) // true ёпта
}
