package glyph

func (pr *Preset) GetType(g string) string {
	if val, ok := pr.GlyphByType[g]; ok {
		return val
	}
	return Undefined
}
