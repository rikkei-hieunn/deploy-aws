package main

import (
	"context"
	"filebus-sample/model"
	"filebus-sample/pkg/locale"
	"filebus-sample/pkg/web"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const location = "Asia/Tokyo"

func init() {
	initLoc()
	initLog()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.Use(
		ErrorHandler(),
		AuthHandler(),
	)
	v1 := router.Group("/v1")
	initPublicRoutes(v1.Group("/public"))

	v2 := router.Group("/file")
	initPublicFollowRoutes(v2)

	log.Info().Msg("server start listening at :8089")
	err := router.Run(":8089")
	if err != nil {
		panic(err)
	}

}

func initPublicRoutes(router *gin.RouterGroup) {
	//ファイル情報
	//router.GET("/me/file", r.publicFileHandler.ListFile)
	router.GET("/me/file/:id/download", DownloadFile)
	//
	////ファイル情報
	//router.POST("/me/file", r.publicFileHandler.UploadFile)
	//router.DELETE("/me/file/:id", r.publicFileHandler.DeleteFile)

	//router.GET("/file/download", DownloadFollowFile)
	//router.POST("/file", r.publicYccFileHandler.UploadFile)
}

func initPublicFollowRoutes(router *gin.RouterGroup) {
	//踏襲API
	//router.GET("/list", r.publicFollowFileHandler.ListFile)
	router.GET("/download", DownloadFollowFile)

	//if !r.config.IsReadonly {
	//	router.POST("/upload", r.publicFollowFileHandler.UploadFile)
	//	router.DELETE("/delete", r.publicFollowFileHandler.DeleteFile)
	//}
}

func initLog() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.LevelFieldName = "level"
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimeFieldFormat = time.RFC3339Nano
}

func initLoc() {
	loc, err := time.LoadLocation(location)
	if err != nil {
		loc = time.FixedZone(location, 9*60*60)
	}
	time.Local = loc
}

// DownloadFile ファイル内容をダウンロード
// @summary ファイル内容をダウンロード
// @tags PublicFile
// @security PublicUser
// @accept json
// @produce json
// @param id path string true "id"
// @success 200 {file} file
// @failure 400 {object} ErrDownloadFileBadRequest "リクエスト内容不正"
// @failure 403 {object} ErrFileForbidden "アクセス禁止"
// @failure 404 {object} ErrFileNotFound "リソースが存在しない"
// @failure 500 {object} ErrInternalError "システムエラー"
// @router /v1/public/me/file/{id}/download [get]
func DownloadFile(c *gin.Context) {
	ctx := c.Request.Context()
	var req model.GetParams
	if !ShouldBindURI(c, &req) {
		return
	}
	err := download(ctx, req.ID, web.NewHTTPWriter(c.Writer))
	if err != nil {
		Error(c, err)
		return
	}

	c.Status(http.StatusOK)
}

func DownloadFollowFile(c *gin.Context) {
	ctx := c.Request.Context()
	var req model.FollowFileParams
	if !ShouldBindQuery(c, &req) {
		return
	}

	req.Path = AppendSlashSuffix(req.Path)

	err := downloadFollow(ctx, req.Path+req.Filename, web.NewHTTPWriter(c.Writer))
	if err != nil {
		Error(c, err)
		return
	}

	c.Status(http.StatusOK)
}

//AppendSlashSuffix 文字列最後に「/」が無い場合追加する
func AppendSlashSuffix(val string) string {
	if len(val) == 0 {
		return ""
	}
	if strings.HasSuffix(val, "/") {
		return val
	}
	return val + "/"
}

func downloadFollow(ctx context.Context, id string, writer *web.HTTPWriter) (err error) {
	writer.WriteFileHeader(id)

	data, err := os.ReadFile(id)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return
}

func download(ctx context.Context, id string, writer *web.HTTPWriter) (err error) {
	writer.WriteFileHeader(id)

	files, err := ioutil.ReadDir("sample-file")
	if err != nil {
		return err
	}

	filePath := "sample-file/"
	for _, f := range files {
		if f.Name() == id {
			filePath += f.Name()

			break
		}
	}

	if filePath == "sample-file/" {
		return fmt.Errorf("file not found")
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	_, err = writer.Write(data)
	if err != nil {
		return err
	}

	return
}

//ShouldBindURI リクエストURLパラメーターをstructに変換する
func ShouldBindURI(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindUri(req); err != nil {
		_ = c.Error(err)
		return false
	}
	return true
}

//ShouldBindQuery リクエストQueryデーターをstructに変換する
func ShouldBindQuery(c *gin.Context, req interface{}) bool {
	if err := c.ShouldBindQuery(req); err != nil {
		_ = c.Error(err)
		return false
	}
	return true
}

func Error(c *gin.Context, err error, meta ...interface{}) {
	ctxErr := c.Error(err)
	if len(meta) > 0 && meta[0] != nil {
		_ = ctxErr.SetMeta(meta)
	}
}

//AuthHandler 認証ハンドラ
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		if strings.HasPrefix(path, "/v1/public/auth") ||
			strings.HasPrefix(path, "/v1/admin/auth") ||
			strings.HasPrefix(path, "/v1/public/user_request") {
			c.Next()
			return
		}

		var username, password string
		var ok bool

		if strings.HasPrefix(path, "/file/") {
			if c.Request.Method == "POST" {
				username = c.PostForm("id")
				password = c.PostForm("pw")
			} else {
				username = c.Query("id")
				password = c.Query("pw")
			}

			if len(username) == 0 || len(password) == 0 {
				_ = c.Error(Required.WithParams("ID and PW"))
				c.Abort()
				return
			}

			//ユーザー・パスワード有効性チェック
			if username != "admin" || password != "admin" {
				_ = c.Error(Required.WithParams("ID and PW"))
				c.Abort()
				return
			}

		} else if path == "/v1/admin/file/file_ycc" {

		} else {
			//踏襲API以外はBasic認証を利用
			username, password, ok = c.Request.BasicAuth()
			if !ok || len(username) == 0 || len(password) == 0 {
				_ = c.Error(AuthAuthenticationHeaderRequired)
				c.Abort()
				return
			}
			if strings.HasPrefix(path, "/v1/admin") {
				//管理者権限が必要
				if username != "admin" || password != "admin" {
					_ = c.Error(AuthAuthenticationHeaderRequired)
					c.Abort()
					return
				}
			} else {
				//ユーザー・パスワード有効性チェック
				if username != "admin" || password != "admin" {
					_ = c.Error(AuthAuthenticationHeaderRequired)
					c.Abort()
					return
				}
			}
		}
		ctx := context.WithValue(c.Request.Context(), username, password)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

var Required = NewCode(http.StatusBadRequest, 13000000)
var AuthAuthenticationHeaderRequired = NewCode(http.StatusUnauthorized, 16000000)
var Internal = NewCode(http.StatusInternalServerError, 11000000)
var codes = make(map[int]struct{})

//Code エラーコード
type Code struct {
	error
	statusCode int
	errorCode  int
	params     []interface{}
}

//NewCode エラーコード初期化
func NewCode(statusCode, errorCode int) Code {
	if _, ok := codes[errorCode]; ok {
		panic(fmt.Sprintf("duplicate error code:%v", errorCode))
	}
	codes[errorCode] = struct{}{}
	return Code{
		statusCode: statusCode,
		errorCode:  errorCode,
	}
}

//StatusCode http status codeを返却
func (c Code) StatusCode() int {
	return c.statusCode
}

//ErrorCode エラーコードを返却
func (c Code) ErrorCode() string {
	return "E" + strconv.Itoa(c.errorCode)
}

//LocalizeKey ローカライズ用のメッセージID
func (c Code) LocalizeKey() string {
	return c.String()
}

//WithParams パラメーター付きでエラー返却
func (c Code) WithParams(params ...interface{}) Code {
	return Code{statusCode: c.statusCode, errorCode: c.errorCode, params: params}
}

//String エラーコードを返却
func (c Code) String() string {
	return "E" + strconv.Itoa(c.errorCode)
}

//Localize エラーコードをメッセージに変換する
func (c Code) Localize(localizer locale.Localizer) string {
	params := make(map[string]interface{}, len(c.params))
	for i, p := range c.params {
		key := "Param" + strconv.Itoa(i)
		params[key] = p
	}

	localized, err := localizer.Localize(c.LocalizeKey(), params)
	if err != nil {
		return c.LocalizeKey()
	}
	return localized
}

func CauseCode(err error) error {
	type causer interface {
		CauseCode() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.CauseCode()
	}
	return err
}

//ErrorHandler エラーハンドリング
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		errorToPrint := c.Errors.ByType(gin.ErrorTypePrivate).Last()
		if errorToPrint == nil {
			return
		}

		ginErr := errorToPrint.Err
		var err Code
		switch v := CauseCode(ginErr).(type) {
		case Code:
			err = v
		default:
			err = Internal
		}
		response := map[string]interface{}{
			"code": err.ErrorCode(),
		}
		if errorToPrint.Meta != nil {
			response["meta"] = errorToPrint.Meta
		}
		//クライアント側にエラー出力
		c.JSON(err.StatusCode(), response)
	}
}
