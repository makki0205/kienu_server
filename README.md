# サーバー仕様
## ファイルupload
### Request
- method
	- POST
- URL
	- makki0250.com:3000/api
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
	- makki0250.com:3000/{download_url}

### 例
`curl http://makki0250.com:3000/file/oI07/ba-ka.txt`

