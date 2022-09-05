package model

//GetParams 取得用パラメーター
type GetParams struct {
	ID string `uri:"id" binding:"required"`
}

//FollowFileParams 踏襲APIファイルパラメーター
type FollowFileParams struct {
	ID       string `form:"id" binding:"required,max=36"`     //ユーザID
	PW       string `form:"pw" binding:"required,max=36"`     //パスワード
	Path     string `form:"path" binding:"required,max=1023"` //ディレクトリパス
	Filename string `form:"file" binding:"required,max=100"`  //ファイル名
	//制御ファイル利用フラグ
	// * 0：制御ファイルを確認しない [デフォルト]
	// * 1：制御ファイルを確認する
	//    * 制御ファイルが無い場合はダウンロードしない。
	Ctlfileflag int  `form:"ctlfileflag" binding:"oneof=0 1"`
	TestFlag    bool `form:"testFlag"` //テスト用フラグ
}
