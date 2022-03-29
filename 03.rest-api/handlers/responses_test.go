package handlers

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/url"
	"os"
	"rest-api/user"
	"strconv"
	"testing"
)

const (
	dbPath = "users.db"
)

type response struct {
	header http.Header
	code   int
	body   []byte
}

type mockWriter response

//constructor
func newMockWriter() *mockWriter {
	return &mockWriter{
		body:   []byte{},
		header: http.Header{},
	}
}

func (mw *mockWriter) Write(b []byte) (int, error) {
	mw.body = make([]byte, len(b))
	for k, v := range b {
		mw.body[k] = v
	}
	return len(b), nil
}

func (mw *mockWriter) WriteHeader(code int) {
	mw.code = code
}

func (mw *mockWriter) Header() http.Header {
	return mw.header
}

func createUserBenchmark(id bson.ObjectId, name string, role string, b *testing.B) *user.User {
	b.StopTimer()
	u := &user.User{
		ID:   id,
		Name: name,
		Role: role,
	}
	b.StartTimer()

	err := u.Save()
	if err != nil {
		b.Fatalf("Error creating a new user: %s", err)
	}
	return u
}

func TestMain(m *testing.M) {
	m.Run()
	os.Remove(dbPath)
}

func prepDb(n int) error {
	os.Remove(dbPath)
	for i := 0; i < n; i++ {
		u := &user.User{
			ID:   bson.NewObjectId(),
			Name: "James_" + strconv.Itoa(i),
			Role: "Mechanic",
		}
		err := u.Save()
		if err != nil {
			return err
		}
	}
	return nil
}

func makeRequest() (*http.Request, error) {
	u, err := url.Parse("/users")
	if err != nil {
		return nil, err
	}
	return &http.Request{
		URL:    u,
		Header: http.Header{},
		Method: http.MethodGet,
	}, nil
}

func getAll(b *testing.B, r *http.Request) {
	prepDb(50)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		mw := newMockWriter()
		b.StartTimer()
		UsersRouter(mw, r)
	}
}

func BenchmarkGetAllNonCached(b *testing.B) {
	r, err := makeRequest()
	if err != nil {
		b.Fatal(err)
	}
	r.Header.Add("Cache-Control", "no-cache")
	getAll(b, r)
}

func BenchmarkGetAllCached(b *testing.B) {
	r, err := makeRequest()
	if err != nil {
		b.Fatal(err)
	}
	getAll(b, r)
}
