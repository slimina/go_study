package memory

import (
	"../../../session"
	"container/list"
	"sync"
	"time"
)

//基于内存的session存储接口的实现

type Provider struct {
	lock     sync.Mutex
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做gc
}

type SessionStore struct {
	sid      string
	lastTime time.Time //最后访问时间
	value    map[interface{}]interface{}
}

var provder = &Provider{list: list.New()}

func (p *Provider) SessionInit(sid string) (session.Session, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newss := &SessionStore{sid: sid, lastTime: time.Now(), value: v}
	el := p.list.PushFront(newss)
	p.sessions[sid] = el
	return newss, nil
}
func (p *Provider) SessionRead(sid string) (session.Session, error) {
	if el, ok := p.sessions[sid]; ok {
		return el.Value.(*SessionStore), nil
	} else {
		sess, err := p.SessionInit(sid)
		return sess, err
	}
	return nil, nil
}

func (p *Provider) SessionUpdate(sid string) error {
	p.lock.Lock()
	defer p.lock.Unlock()
	if el, ok := p.sessions[sid]; ok {
		el.Value.(*SessionStore).lastTime = time.Now()
		p.list.MoveToFront(el)
	}
	return nil
}

func (p *Provider) SessionDestory(sid string) error {
	if el, ok := p.sessions[sid]; ok {
		delete(p.sessions, sid)
		p.list.Remove(el)
	}
	return nil
}

func (p *Provider) SessionGC(maxLiveTime int64) {
	p.lock.Lock()
	defer p.lock.Unlock()
	for {
		el := p.list.Back()
		if el == nil {
			break
		}
		if el.Value.(*SessionStore).lastTime.Unix()+maxLiveTime < time.Now().Unix() {
			p.list.Remove(el)
			delete(p.sessions, el.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (s *SessionStore) Set(key, value interface{}) error {
	s.value[key] = value
	provder.SessionUpdate(s.sid)
	return nil
}

func (s *SessionStore) Get(key interface{}) interface{} {
	provder.SessionUpdate(s.sid)
	if val, ok := s.value[key]; ok {
		return val
	}
	return nil
}

func (s *SessionStore) Delete(key interface{}) error {
	delete(s.value, key)
	provder.SessionUpdate(s.sid)
	return nil
}

func (s *SessionStore) SessionID() string {
	return s.sid
}

func init() {
	provder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", provder)
}
