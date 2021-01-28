package main

type AnchorEntry struct {
	DBHeight uint32 `json:"dbheight"`
	KeyMR    string `json:"keymr"`
}

type AnchorProof struct {
	AnchorRecordVer string `json:"anchorrecordver"`
	DBHeight        uint32 `json:"dbheight"`
	KeyMR           string `json:"keymr"`
	Factom          struct {
		TXID        string `json:"txid"`
		BlockHeight uint32 `json:"blockheight"`
		BlockHash   string `json:"blockhash"`
	}
}
