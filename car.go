package car

// Car structure
type Car struct {
	Mark                string
	Year, Vol, Trunkvol int
	Start, Openw        bool
}

// Procvol - procent
func Procvol(trunkvol, vol int) int {
	return trunkvol * 100 / vol
}

// Openwin - translate
func Openwin(window bool) string {
	var openw string
	if window == false {
		openw = "Закрыто"
	} else {
		openw = "Открыто"
	}
	return openw
}

// Enginestart - translate
func Enginestart(engine bool) string {
	var sengine string
	if engine == false {
		sengine = "Не заведен"
	} else {
		sengine = "Заведен"
	}
	return sengine
}

// Сartrunk - trunk and car
func Сartrunk(vol int) string {
	var volcar string
	if vol > 700 {
		volcar = "Грузовой автомобиль"
	} else {
		volcar = "Легковой автомобиль"
	}
	return volcar
}

//GetVal - print val
func (trunk Car) GetVal() (string, string, int, int, int, string, string, string) {
	return Сartrunk(trunk.Vol), trunk.Mark, trunk.Year, trunk.Vol, trunk.Trunkvol, "%", Enginestart(trunk.Start), Openwin(trunk.Openw)
}

//SetVal - set val
func (trunk *Car) SetVal(Mark string, Year, Vol, TrunkVol int, Start, Openw bool) {
	trunk.Mark = Mark
	trunk.Year = Year
	trunk.Vol = Vol
	trunk.Trunkvol = TrunkVol * 100 / Vol
	trunk.Start = Start
	trunk.Openw = Openw
}
