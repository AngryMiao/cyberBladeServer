package keyboardSerializer

type ChangeData struct {
	KeyCode  string `json:"key_code"`
	KeyLevel string `json:"key_level"`
	Row      string `json:"row"`
	Col      string `json:"col"`
}

type EmbeddingDbQuery struct {
	IsVia    bool     `json:"isVia"`
	Keyboard string   `json:"keyboard"`
	KeyLevel int      `json:"key_level"`
	Text     []string `json:"text"`
}

type Change80KeyMap struct {
	MsgType  int    `json:"msg_type"`  // 0: '实时改键', 1: 'press&hold', 2: '宏',
	LayerNum int    `json:"layer_num"` //当前键层
	RowNum   int    `json:"row_num"`   //行 Math.floor(index / 25)
	Index    int    `json:"index"`     //列 index % 25
	OldKey   string `json:"old_key"`   //旧键值 oldKey
	NewKey   string `json:"new_key"`   //新键值 key
}

type CbKeyMaps struct {
	Description []struct {
		Text string `json:"text"`
	} `json:"description"`
	KeyCode string `json:"key_code"`
}
