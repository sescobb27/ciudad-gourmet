package session

// Provider contains global session methods and saved SessionStores.
// it can operate a SessionStore by its id.
// Based on from github.com/astaxie/beego/session
type Provider interface {
    SessionStore(sessionId string) SessionStore
    SessionDestroy(sessionId string)
    SessionRead(sessionId string) SessionStore
}

type MemoryProvider struct {
    Provider
    sessions  map[string]SessionStore // map in memory
    semaphore chan signal
}

func NewSessionProvider() *MemoryProvider {
    mp := &MemoryProvider{
        sessions:  make(map[string]SessionStore),
        semaphore: make(chan signal, 1),
    }
    mp.semaphore <- signal{}
    return mp
}

func (mp *MemoryProvider) SessionStore(sessionId string) SessionStore {
    <-mp.semaphore
    sessionStore, exist := mp.sessions[sessionId]
    if !exist {
        sessionStore = NewMemorySessionStore(sessionId)
        mp.sessions[sessionId] = sessionStore
    }
    mp.semaphore <- signal{}
    return sessionStore
}

func (mp *MemoryProvider) SessionRead(sessionId string) SessionStore {
    <-mp.semaphore
    sessionStore, exist := mp.sessions[sessionId]
    mp.semaphore <- signal{}
    if exist {
        return sessionStore
    }
    return nil
}

func (mp *MemoryProvider) SessionDestroy(sessionId string) {
    <-mp.semaphore
    if _, exist := mp.sessions[sessionId]; exist {
        delete(mp.sessions, sessionId)
    }
    mp.semaphore <- signal{}

}
