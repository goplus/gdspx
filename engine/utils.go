package engine

import (
	"grow.graphics/gd"
	_ "grow.graphics/gd/gdextension"
)

var (
	rootSelf  *EngineNode
	Root      gd.Node2D
	KeepAlive gd.Lifetime
	Temporary gd.Lifetime

	AudioUtil     *AudioMgr
	AnimationUtil *AnimationMgr
	PhysicUtil    *PhysicMgr
	RenderUtil    *RenderMgr
	InputUtil     *InputMgr
)

func GetTree() gd.SceneTree {
	return rootSelf.Super().AsNode().GetTree(Temporary)
}
func ToString(s string) gd.String {
	return Temporary.String(s)
}
func ToNodePath(s string) gd.NodePath {
	return ToString(s).NodePath(Temporary)
}
func ToStringName(s string) gd.StringName {
	return Temporary.StringName(s)
}
func ToVariant(v any) gd.Variant {
	return Temporary.Variant(v)
}
func TweenVector2(pself gd.Object, tween gd.Tween, property string, final_val gd.Vector2, duration float32) gd.PropertyTweener {
	return tween.TweenProperty(
		Temporary, pself,
		Temporary.String(property).NodePath(Temporary),
		Temporary.Variant(final_val),
		gd.Float(duration))
}
func DoTweenVector2(pself gd.Object, property string, dstValue gd.Vector2, duration float32, fn any) {
	tween := GetTree().CreateTween(Temporary)
	TweenVector2(pself, tween, property, dstValue, duration)
	tween.TweenCallback(Temporary, Temporary.Callable(fn))
}
func DoTweenChainVector2(pself gd.Object, property string, dstValue gd.Vector2, duration float32, chainVal gd.Vector2, cahainDuration float32, fn any) {
	tween := GetTree().CreateTween(Temporary)
	TweenVector2(pself, tween, property, dstValue, duration)
	TweenVector2(pself, tween.Chain(Temporary), property, chainVal, 4)
	tween.TweenCallback(Temporary, Temporary.Callable(fn))
}
