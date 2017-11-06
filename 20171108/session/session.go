package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//session接口定义
type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionID() string
}

//存储提供者接口
type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestory(sid string) error
	SessionGC(maxLiveTime int64)
}

//定义存储集合map
var providers = make(map[string]Provider)

//注册存储实例
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is null")
	}
	//判断name是否存在
	if _, ok := providers[name]; ok {
		panic("session: Register called twice for provide " + name)
	}
	providers[name] = provider
}

//cookie manager
type Manager struct {
	cookieName  string     //private
	lock        sync.Mutex //锁
	provider    Provider
	maxLiveTime int64
}

//初始化 manager
func NewManager(provideName, cookieName string, maxLiveTime int64) (*Manager, error) {
	provider, ok := providers[provideName]
	if !ok {
		return nil, fmt.Errorf("session : unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{cookieName: cookieName, provider: provider, maxLiveTime: maxLiveTime}, nil
}

func (m *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//获取session
func (m *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	m.lock.Lock()
	defer m.lock.Unlock()
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		sid := m.sessionId()
		session, _ = m.provider.SessionInit(sid)
		//httpOnly 设置是否可通过客户端脚本访问这个设置的cookie,劫持防范
		cook := http.Cookie{Name: m.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(m.maxLiveTime)}
		http.SetCookie(w, &cook)
		fmt.Println("SessionStart:", url.QueryEscape(sid))
	} else {
		fmt.Println("SessionStart get:", cookie.Value)
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = m.provider.SessionRead(sid)
	}
	return
}

//销毁session

func (m *Manager) SessionDestroy(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(m.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	fmt.Println("SessionDestroy:", cookie.Value)
	m.lock.Lock()
	defer m.lock.Unlock()
	sid, _ := url.QueryUnescape(cookie.Value)
	m.provider.SessionDestory(sid)
	exp := time.Now()
	cook := http.Cookie{Name: m.cookieName, Path: "/", HttpOnly: true, Expires: exp, MaxAge: -1}
	http.SetCookie(w, &cook)
}

func (m *Manager) GC() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.provider.SessionGC(m.maxLiveTime)
	time.AfterFunc(time.Duration(m.maxLiveTime)*time.Second, func() {
		m.GC()
	})
}
