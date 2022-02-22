package storage

//fetch dataPage from disk
func FetchDataPage(pageId uint32) *dataPage {
	dp := &dataPage{}
	return dp
}

//fetch logPage from disk
func FetchLogPage(pageId uint32) *logPage {
	lp := &logPage{}
	return lp
}

//swap dataPage into disk
func SwapDataPage(dp *dataPage) {

}

//swap logPage into disk
func SwapLogPage(lp *logPage) {

}
