package types

func (m *RoundMeta) IsInRound(sequence int) bool {
	return sequence >= m.RoundStart && sequence <= m.End
}

func (m *RoundMeta) IsInBuyPhase(sequence int) bool {
	return sequence >= m.RoundStart && sequence < m.PlayStart
}

func (m *RoundMeta) IsInPlayPhase(sequence int) bool {
	return sequence >= m.PlayStart && sequence <= m.End
}
