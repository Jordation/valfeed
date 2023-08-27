package types

func (m *RoundData) IsInRound(sequence int) bool {
	return sequence >= m.RoundStart && sequence <= m.End
}

func (m *RoundData) IsInBuyPhase(sequence int) bool {
	return sequence >= m.RoundStart && sequence < m.PlayStart
}

func (m *RoundData) IsInPlayPhase(sequence int) bool {
	return sequence >= m.PlayStart && sequence <= m.End
}
