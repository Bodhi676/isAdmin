package cache

// gRPC присвоение админки пользователю
func (c *Cache) SetUserToAdmin(id UserID) {
	// Проверяем статус пользователя
	status, ok := c.Get(id.GetUser())
	if !ok {
		// если пользователь отсутствует в kэше cоздаём его
		c.Set(id.GetUser(), true, AdminTime)
	}
	if !status {

	}

}

// gRPC возвращение статуса пользователя
// func GetUserIsAdmin(id cache.UserID) bool {

// }
