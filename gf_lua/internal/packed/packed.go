package packed

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var path = "resource/scripts/"

type scriptEVALSha struct {
	r *gredis.Redis
	m sync.Map
}

var ScriptEVALSha = &scriptEVALSha{}

// 初始化 对象
func new(r *gredis.Redis) *scriptEVALSha {
	if r == nil {
		return nil
	}
	return &scriptEVALSha{
		r: r,
		m: sync.Map{},
	}
}

func init() {
	ctx := context.Background()
	ScriptEVALSha = new(g.Redis())
	list, err := gfile.ScanDirFile(path, "*.lua", true)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, v := range list {
		sha, err := ScriptEVALSha.registerScriptFile(ctx, v)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(sha)
		}
	}
}

// @Description: 获取文件名称
// @param ctx
// @param filePath
// @return fileName
func getNameByFilePath(ctx context.Context, filePath string) (fileName string) {
	//获取文件名
	base := filepath.Base(filePath)
	if i := strings.LastIndexByte(base, '.'); i != -1 {
		fileName = base[:i]
	} else {
		fileName = base
	}
	return
}

// @Description: 注册的lua脚本
// @receiver s
// @param ctx
// @param filePath  路径
// @return interface{}
// @return interface{}
func (dr *scriptEVALSha) registerScriptFile(ctx context.Context, filePath string) (sha string, err error) {
	fileName := getNameByFilePath(ctx, filePath)
	g.Log().Debugf(ctx, "script: fileName = %s ", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		g.Log().Error(ctx, "open file failed:", err, filePath)
		return "", err
	}
	defer file.Close()
	fileBuffer, err := io.ReadAll(file)
	if err != nil {
		g.Log().Error(ctx, "read file failed:", err, filePath)
		return "", err
	}
	script := string(fileBuffer)

	// 缓存脚本,获取校验和
	if sha, err = dr.scriptLoad(ctx, script); err != nil && sha == "" {
		return "", errors.New(fmt.Sprintf("script load `%s` failed:", err))
	}
	g.Log().Debugf(ctx, "script: fileName = %s sha = %s", fileName, sha)

	// 注册脚本key和校验和，如果fileName已注册，则返回错误
	res, loaded := dr.m.LoadOrStore(fileName, sha)
	if loaded {
		err = errors.New(fmt.Sprintf("script fileName = %s has registered. exists sha1 = %s", fileName, res))
		return "", err
	}
	sha = res.(string)
	g.Log().Debugf(ctx, "register script fileName: key = %s, sha = %s", fileName, sha)
	return sha, err
}

// 缓存脚本  sha1加密
func (dr *scriptEVALSha) scriptLoad(ctx context.Context, script string) (string, error) {
	var (
		doArgs = []interface{}{"LOAD", script}
	)
	if v, err := dr.r.Do(ctx, "SCRIPT", doArgs...); err != nil {
		return "", err
	} else {
		return v.String(), nil
	}
}

// @Description:  执行的脚本
// @param ctx
// @param fileName  sha1之后脚本名
// @param keys
// @param argv
// @return *gvar.Var
// @return error
func (dr *scriptEVALSha) evalSha1CmdScriptKey(ctx context.Context, fileName string, keys []string, argv ...interface{}) (*gvar.Var, error) {
	var (
		//第一个参数是执行脚本 第二参数是key的个数
		doArgs = make([]interface{}, 2+len(keys), 2+len(keys)+len(argv))
	)
	g.Log().Debugf(ctx, "keys: %v, args: %v", keys, argv)
	// 读取脚本对应的SHA1校验和
	sha, ok := dr.m.Load(fileName)
	if !ok {
		return nil, errors.New(fmt.Sprintf("fileName = %s have not registered", fileName))
	}
	doArgs[0] = sha
	doArgs[1] = len(keys)

	for i, k := range keys {
		doArgs[2+i] = k
	}
	doArgs = append(doArgs, argv...)
	g.Log().Debugf(ctx, "eval script key: `evalsha %v`", doArgs)
	return dr.r.Do(ctx, "evalsha", doArgs...)
}

// @Description: 清空脚本
// @param ctx
// @return string
// @return error
func (dr *scriptEVALSha) FlushScriptKey(ctx context.Context) (string, error) {
	var (
		doArgs = []interface{}{"FLUSH"}
		newMp  = sync.Map{}
	)
	if v, err := dr.r.Do(ctx, "SCRIPT", doArgs...); err != nil {
		return "", err
	} else {
		dr.m = newMp
		return v.String(), nil
	}
}

func (dr *scriptEVALSha) GetAndDel(ctx context.Context, key string) (*gvar.Var, error) {
	return dr.evalSha1CmdScriptKey(ctx, "get-and-del", []string{key})
}

func (dr *scriptEVALSha) GetAndSet(ctx context.Context, key string, v ...interface{}) (*gvar.Var, error) {
	return dr.evalSha1CmdScriptKey(ctx, "get-and-set", []string{key}, v)
}

// Map 获取所有键值对
func (dr *scriptEVALSha) Map() map[string]string {
	var m = make(map[string]string)
	dr.m.Range(func(key, value interface{}) bool {
		if k, ok := key.(string); ok {
			m[k] = value.(string)
		}
		return true
	})

	return m
}

// Keys 获取所有key
func (dr *scriptEVALSha) Keys() []string {
	var keys []string
	dr.m.Range(func(key, _ interface{}) bool {
		if k, ok := key.(string); ok {
			keys = append(keys, k)
		}
		return true
	})

	return keys
}
