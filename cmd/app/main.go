package main

import (
	"fmt"
	"time"

	"github.com/Bodhi676/isAdmin/internal/cache"
)

const (
	AdminTime      = 5 * time.Minute
	ClearCacheTime = 10 * time.Minute
)

func main() {
	// Тестируем работу кэша
	// Создаем контейнер с временем жизни по-умолчанию равным 5 минут(AdminTime) и удалением просроченного кеша каждые 10 минут(ClearCacheTime)
	Mycache := cache.New(AdminTime, ClearCacheTime)

	// Установить кеш с ключем 365557(uint64) и временем жизни 5 минут
	usr := cache.NewUser(365557)
	Mycache.Set(usr.GetUser(), true, AdminTime)

	// Получить кеш с ключем 365557
	i, _ := Mycache.Get(usr.GetUser())

	fmt.Println(i) // true ёпта

}
