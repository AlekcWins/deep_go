package structs

type Stats struct {
	// respectAndStrength (1 байт = 8 бит)
	// 7 6 5 4 | 3 2 1 0
	// --------+--------
	// respect | strength
	respectAndStrength TwoNumberByte // 1   //max value need ([0…10]) for each [RRRR][SSSS]

	// experienceAndLevel (1 байт = 8 бит)
	// 7 6 5 4 | 3 2 1 0
	// --------+--------
	// exp     | level
	experienceAndLevel TwoNumberByte // 1   //max value need ([0…10]) for each [EEEE][LLLL]

}

func (s *Stats) GetRespect() int32 {
	return int32(s.respectAndStrength.GetHighValue())
}

func (s *Stats) SetRespect(respect int32) bool {
	if respect > 10 {
		return false
	}
	return s.respectAndStrength.SetHighValue(byte(respect))
}

func (s *Stats) GetStrength() int32 {
	return int32(s.respectAndStrength.GetLowValue())
}

func (s *Stats) SetStrength(strength int32) bool {
	return s.respectAndStrength.SetLowValue(byte(strength), 10)
}

func (s *Stats) GetExperience() int32 {
	return int32(s.experienceAndLevel.GetHighValue())
}

func (s *Stats) SetExperience(experience int32) bool {
	if experience > 10 {
		return false
	}
	return s.experienceAndLevel.SetHighValue(byte(experience))
}

func (s *Stats) GetLevel() int32 {
	return int32(s.experienceAndLevel.GetLowValue())
}

func (s *Stats) SetLevel(level int32) bool {
	return s.experienceAndLevel.SetLowValue(byte(level), 10)
}
