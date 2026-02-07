package jobs
import "sync"

/*
shared state
map concurrency problemi
mutex vs RWMutex
pointer state
atomic update pattern
callback update pattern
*/

//thread-safe job storage
type Store struct {
	mu sync.RWMutex //more than one goroutine could read at the same time, as long as they're not writing
	jobs map[string]*JobState //Shared state’e erişen HER yer mutex kullanmalıdır.
}

//başlatmamız lazım çünkü mapin initialize edilmesi gerekiyor yoksa panic verir
func NewStore () *Store {
	return &Store{ 
		jobs: make(map[string]*JobState), //job id is the key
	}
}

//returns the job with given id, if not found then returns false
/* 
aynı anda birçok reader olabilir
writer varken reader olamaz
reader varken writer bekler
*/
func (s *Store) Get (id string) (*JobState, bool) {
	s.mu.RLock() //State değiştirmiyoruz.
	defer s.mu.RUnlock()
	job , ok := s.jobs[id]
	return job, ok

}

//will put the id + Jobstate into the map
func (s *Store) Set (id string, state *JobState) {
	if id == "" {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.jobs[id] = state

}

//find the given job, update it safely ->atomic update
func (s *Store) Update (id string, fn func(*JobState)) {
	if id == "" {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	job, ok := s.jobs[id]
	if !ok {
		return
	}
	fn(job)
}