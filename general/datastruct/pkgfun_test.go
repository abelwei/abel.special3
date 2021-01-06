package datastruct

import "testing"

func TestSettingTxt2Map(t *testing.T) {
	txt := `[base_nc]
host="111W111111"
port="222"`
	t.Log(SettingTxt2Map(txt, "toml"))
}

func TestYml2MapSlice(t *testing.T) {
	txt := `
- string: name
- string: int1
- float: cost
- json:
    - string: name
`
t.Log(Yml2MapSlice(txt))
}