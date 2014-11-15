package session

// semaphores implemented as https://groups.google.com/forum/#!msg/golang-dev/ShqsqvCzkWg/Kg30VPN4QmUJ
// in resume: We want the property that an Unlock happens before the next Lock,
//            and there is no value being communicated between those two operations.
//            Unlock is moving one value out of the buffer to make room for a different value in Lock.
//            The two are, strictly speaking, unrelated operations.

type signal struct{}

// Based on from github.com/astaxie/beego/session
type SessionStore interface {
    Set(key string, value interface{}) //set session value
    Get(key string) interface{}        //get session value
    Delete(key string)                 //delete session value
    SessionID() string                 //back current sessionID
    Flush()                            //delete all data
}

// memory session store.
// it saved sessions in a map in memory.
// based on from github.com/astaxie/beego/session/sess_mem.go
type MemorySessionStore struct {
    SessionStore                        //implements SessionStore interface
    sessionId    string                 //session id
    session      map[string]interface{} //session store
    semaphore    chan signal            //semaphore
}

func NewMemorySessionStore(sessionId string) *MemorySessionStore {
    mStore := &MemorySessionStore{
        sessionId: sessionId,
        session:   make(map[string]interface{}),
        semaphore: make(chan signal, 1),
    }
    mStore.semaphore <- signal{}
    return mStore
}

// set user to memory session
func (st *MemorySessionStore) Set(key string, value interface{}) {
    <-st.semaphore
    st.session[key] = value
    st.semaphore <- signal{}
}

// get user from memory session by key
func (st *MemorySessionStore) Get(key string) interface{} {
    <-st.semaphore
    user, ok := st.session[key]
    st.semaphore <- signal{}
    if ok {
        return user
    } else {
        return nil
    }
}

// delete in memory session by key
func (st *MemorySessionStore) Delete(key string) {
    <-st.semaphore
    delete(st.session, key)
    st.semaphore <- signal{}
}

// clear all users in memory session
func (st *MemorySessionStore) Flush() {
    <-st.semaphore
    st.session = make(map[string]interface{})
    st.semaphore <- signal{}
}

// get this id of memory session store
func (st *MemorySessionStore) SessionID() string {
    return st.sessionId
}
