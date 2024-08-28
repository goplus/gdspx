package sprite

import (
	. "godot-ext/gdspx/pkg/engine"
)

type Player struct {
	Actor
	OnDieEvent *Event0

	runSpeedDamping float32
	PlayerMode      EPlayerMode
	isControlled    bool
}

type EPlayerMode int64

const (
	SMALL EPlayerMode = iota
	BIG
	SHOOTING
)

var PlayModeAnimName = [...]string{
	"small",
	"big",
	"shooting",
}

var (
	SPEED         = float32(600)
	JUMP_VELOCITY = float32(400.0)
	PLAYER_MODE   = int64(1)
	GRAVITY       = float32(-980)

	minStompDegree = float32(25)
	maxStompDegree = float32(155)

	ShouldSyncCamera = true
)

func (pself *Player) OnStart() {
	pself.OnDieEvent = NewEvent0()
	pself.runSpeedDamping = 0.5
	pself.PlayerMode = SMALL
	pself.isControlled = false
}
func (pself *Player) OnUpdate(delta float32) {
	if pself.GetPosX() > CameraMgr.GetCameraPosition().X && ShouldSyncCamera {
		CameraMgr.SetCameraPosition(Vec2{pself.GetPosX(), CameraMgr.GetCameraPosition().Y})
	}
}

func (pself *Player) OnFixedUpdate(delta float32) {
	if pself.isControlled {
		return
	}
	cameraRect := CameraMgr.GetViewportRect()
	cameraLeftBound := CameraMgr.GetCameraPosition().X - cameraRect.Size.X/2/CameraMgr.GetCameraZoom().X

	if !pself.IsOnFloor() {
		pself.AddVelY(GRAVITY * delta)
	}

	if pself.GetPosX() < cameraLeftBound+8 && pself.GetVelX() < 0 {
		pself.SetVelocity(Vec2{0, 0})
		return
	}

	if InputMgr.IsActionJustPressed("jump") && pself.IsOnFloor() {
		pself.SetVelY(JUMP_VELOCITY)
	}
	if InputMgr.IsActionJustReleased("jump") && pself.GetVelY() > 0 {
		pself.SetVelY(pself.GetVelY() * 0.5)
	}

	dir := InputMgr.GetAxis("left", "right")
	if dir != 0 {
		pself.SetVelX(Lerpf(pself.GetVelX(), dir*SPEED, pself.runSpeedDamping*delta))
	} else {
		pself.SetVelX(MoveToward(pself.GetVelX(), 0, SPEED))
	}

	if InputMgr.IsActionJustPressed("shoot") && pself.PlayerMode == SHOOTING {
		pself.shoot()
	} else {
		pself.TriggeraAnimation(pself.GetVelocity(), dir, pself.PlayerMode)
	}

	pself.MoveAndSlide()
}
func (pself *Player) setAnimSign(isRight bool) {
	dir := int64(1)
	if !isRight {
		dir = -1
	}
	pself.SetChildScale("AnimatedSprite2D", Vec2{float32(dir), 1})
}

func (pself *Player) getAnimDir() int64 {
	scale := pself.GetChildScale("AnimatedSprite2D").X
	if scale < 0 {
		return -1
	}
	return 1
}
func (pself *Player) TriggeraAnimation(velocity Vec2, direction float32, playerMode EPlayerMode) {
	animation_prefix := PlayModeAnimName[int64(playerMode)]
	if !pself.IsOnFloor() {
		pself.PlayAnimation(animation_prefix + "_jump")
	} else if Sign(velocity.X) != Sign(direction) && velocity.X != 0 && direction != 0 {
		pself.PlayAnimation(animation_prefix + "_slide")
		pself.setAnimSign(direction > 0.5)
	} else {
		if pself.getAnimDir() != Sign(velocity.X) && Sign(velocity.X) != 0 {
			pself.setAnimSign(Sign(velocity.X) == 1)
		}
		if pself.GetVelX() != 0 {
			pself.PlayAnimation(animation_prefix + "_run")
		} else {
			pself.PlayAnimation(animation_prefix + "_idle")
		}
	}
}

func (pself *Player) shoot() {
	bullet := CreateSprite[Bullet]()
	pos := pself.GetPosition()
	bullet.SetPosition(Vec2{pos.X, pos.Y})
}
func (pself *Player) OnHit() {
	pself.OnDieEvent.Trigger()
}

func (pself *Player) OnTriggerEnter(target ISpriter) {
	if pself.isControlled {
		return
	}
	if enemy, ok := target.(*Goomba); ok {
		angleOfCollision := RadToDeg(pself.GetPosition().AngleToPoint(enemy.GetPosition()))
		println("angleOfCollision", int(angleOfCollision))
		if angleOfCollision > minStompDegree && maxStompDegree > angleOfCollision {
			enemy.Die()
		} else {
			pself.Die()
		}
	}
}

func (pself *Player) Die() {
	pself.DisablePhysic()
	pself.isControlled = true
	TweenPos2(pself, Vec2{pself.GetPosX(), pself.GetPosY() + 25}, 0.3, Vec2{pself.GetPosX(), pself.GetPosY() - 500}, 2, func() {
		pself.Destroy()
	})
}
