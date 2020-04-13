# logger

## 如何使用
### 導入
```
 go get -u --insecure 10.40.42.38/BP05G0/go-logger
```
### 使用
創建實例

```
logger.Log = logger.NewLogger(logger.Logger{
	Level: "debug",
	Format: "json",
})
//Debug
logger.Log.Deub("test")
```