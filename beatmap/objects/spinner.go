package objects

import (
	"danser/bmath"
	"danser/render"
	"danser/settings"
	"github.com/go-gl/mathgl/mgl32"
	"strconv"
)

type Spinner struct {
	objData *basicData
	pos     bmath.Vector2d
	Timings *Timings
}

func NewSpinner(data []string) *Spinner {
	spinner := &Spinner{}
	spinner.objData = commonParse(data)
	endtime, _ := strconv.ParseInt(data[5], 10, 64)
	spinner.objData.EndTime = int64(endtime)
	spinner.pos = bmath.Vector2d{256,192}
	return spinner
}

func (self Spinner) GetBasicData() *basicData {
	return self.objData
}

func (self *Spinner) SetTiming(timings *Timings) {
	self.Timings = timings
}

func (self *Spinner) GetPosition() bmath.Vector2d {
	return self.pos
}

func (self *Spinner) Update(time int64) bool {
	return true
}

func (self *Spinner) Draw(time int64, preempt float64, color mgl32.Vec4, batch *render.SpriteBatch) bool {

	alpha := 1.0

	if time < self.objData.StartTime {
		return false
	} else if time < self.objData.StartTime + int64(preempt){
		alpha = float64(color[3]) / preempt
	}else {
		alpha = float64(color[3])
	}

	batch.SetTranslation(self.objData.StartPos)

	batch.SetScale(10, 10)
	batch.SetColor(1, 1, 1, alpha)
	batch.DrawUnit(*render.Spinner)

	batch.SetSubScale(1, 1)

	if time >= self.objData.EndTime+int64(preempt/4) {
		return true
	}
	return false
}

func (self *Spinner) SetDifficulty(preempt, fadeIn float64) {

}

func (self *Spinner) DrawApproach(time int64, preempt float64, color mgl32.Vec4, batch *render.SpriteBatch) {
	alpha := 1.0
	// 计算AR
	fake_preempt := float64(self.objData.StartTime - self.objData.EndTime) / settings.General.SpinnerMult
	arr := float64(self.objData.EndTime-time) / fake_preempt

	if time < self.objData.StartTime {
		alpha = 0
	} else if time < self.objData.StartTime + int64(preempt){
		alpha = float64(color[3]) / preempt
	}else {
		alpha = float64(color[3])
	}

	batch.SetTranslation(self.objData.StartPos)

	if time <= self.objData.EndTime {
		batch.SetColor(float64(color[0]), float64(color[1]), float64(color[2]), alpha)
		batch.SetSubScale(arr*2, arr*2)
		batch.DrawUnit(*render.ApproachCircle)
	}

	batch.SetSubScale(0, 0)
}