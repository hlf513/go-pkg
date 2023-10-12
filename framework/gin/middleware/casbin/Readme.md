# Casbin

## 技术方案
- `model.conf` 默认值见 `default.go`
- `policy.csv` 使用的是 `MySQL` 动态调整，驱动使用`gorm`
- 分布式消息的监视器，使用的是`Redis`

## 依赖
### 公共库
- go-redis
- cast
- gin
- gin-jwt

### 内置库
- mysql
- zap
- response
- redis
