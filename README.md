# サーバー仕様
## ファイルupload
### Request
- method
	- POST
- URL
	- http://tmp.fun/api/upload
- parameter
	- file : ファイル

	
### Response

```json
{
  "Description_url": "/api/@oI07",
  "download_url": "/file/oI07/ba-ka.txt",
  "uuid": "oI07"
}
```
## ファイルdownload
### Request
- method
	- GET
- URL
	- http://tmp.fun/{download_url}

### 例
`curl http://tmp.fun/file/oI07/ba-ka.txt`


## ファイルの詳細情報取得
- method
	- GET
- URL
	- http://tmp.fun/api/@{uuid}
### 例
```json
{
  "create_at": "2017-07-02T04:43:24+09:00",
  "download_url": "/file/oI07/ba-ka.txt",
  "exp_at": "2017-07-09T04:43:24+09:00",
  "file_size": "74.5K"
}
```