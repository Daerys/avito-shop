package entity

type CoinTransaction struct {
	ID         int64 `db:"id"`
	FromUserID int64 `db:"from_user_id"`
	ToUserID   int64 `db:"to_user_id"`
	Amount     int   `db:"amount"`
}
