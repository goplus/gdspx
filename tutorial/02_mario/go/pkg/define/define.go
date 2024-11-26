package define

var (
	WinWidth        = float32(1080)
	WinHeight       = float32(1920)
	Score     int64 = 0
)

const (
	CollisionLayer_Player = int64(1<<1) >> 1
	CollisionLayer_Enemy  = int64(1<<3) >> 1
	CollisionLayer_Pipe   = int64(1<<5) >> 1
	CollisionLayer_Bullet = int64(1<<7) >> 1
)
