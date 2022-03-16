package user

import (
	"github.com/asdine/storm/v3"
	"gopkg.in/mgo.v2/bson"
	"os"
	"reflect"
	"strconv"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
	os.Remove(dbPath)
}

func compareUserObjs(t *testing.T, u1 *User, u2 *User) {
	if !reflect.DeepEqual(u1, u2) {
		t.Error("Records are not matching!")
	}
}

func createUser(id bson.ObjectId, name string, role string, t *testing.T) *User {
	u := &User{
		ID:   id,
		Name: name,
		Role: role,
	}
	err := u.Save()
	if err != nil {
		t.Fatalf("Error creating a new user: %s", err)
	}
	return u
}

func createUserBenchmark(id bson.ObjectId, name string, role string, b *testing.B) *User {
	b.StopTimer()
	u := &User{
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

func TestCRUD(t *testing.T) {
	// Create Record Test
	t.Log("Create")
	u1 := createUser(
		bson.NewObjectId(),
		"Peter",
		"Creative",
		t,
	)

	// Read Record Test
	t.Log("Read")
	u1Read, err := One(u1.ID)
	if err != nil {
		t.Fatalf("Error retrieving an existing user: %s", err)
	}
	compareUserObjs(t, u1, u1Read)

	// Update Record Test
	t.Log("Update")
	u1.Role = "Documents engineering"
	err = u1.Save()
	if err != nil {
		t.Fatalf("Error saving a record: %s", err)
	}
	u1Read, err = One(u1.ID)
	if err != nil {
		t.Fatalf("Error retrieving an updated user: %s", err)
	}
	compareUserObjs(t, u1Read, u1)

	// Delete Record Test
	t.Log("Delete")
	err = Delete(u1.ID)
	if err != nil {
		t.Fatalf("Error removing user record: %s", err)
	}

	//	Retrieving Non-existent Record
	t.Log("Non Existent")
	u1, err = One(u1.ID)
	if err == nil {
		t.Fatal("Record should not exist anymore")
	}
	if err != storm.ErrNotFound {
		t.Fatalf("Error retriving a nonexistent record: %s", err)
	}

	//	Retrieving All Users
	t.Log("Read all users")
	createUser(bson.NewObjectId(), "Alberto", "CTO", t)
	createUser(bson.NewObjectId(), "Marvin", "Chief Engineer", t)
	createUser(bson.NewObjectId(), "Marvin", "Chief Engineer", t)
	var users []User
	users, err = All()
	if err != nil {
		t.Fatalf("Error reading all records: %s", err)
	}

	if len(users) != 3 {
		t.Errorf("Unexpected number of users retrieved. Expected: 3 / Actual: %d", len(users))
	}
}

func cleanDb(b *testing.B) {
	os.Remove(dbPath)
	createUserBenchmark(
		bson.NewObjectId(),
		"Peter",
		"Creative",
		b,
	)
	b.ResetTimer()
}

func BenchmarkCreate(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		// Create Record Test
		createUserBenchmark(
			bson.NewObjectId(),
			"Peter_"+strconv.Itoa(i),
			"Creative",
			b,
		)
	}
}

func BenchmarkRead(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := createUserBenchmark(
			bson.NewObjectId(),
			"Peter_"+strconv.Itoa(i),
			"Creative",
			b,
		)
		b.StartTimer()
		// Read Record Test
		_, err := One(u.ID)
		if err != nil {
			b.Fatalf("Error retrieving an existing user: %s", err)
		}
	}
}

func BenchmarkUpdate(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := createUserBenchmark(
			bson.NewObjectId(),
			"Peter_"+strconv.Itoa(i),
			"Creative",
			b,
		)
		b.StartTimer()
		// Update Record Test
		u.Role = "Documents engineering"
		err := u.Save()
		if err != nil {
			b.Fatalf("Error saving a record: %s", err)
		}
	}
}

func BenchmarkDelete(b *testing.B) {
	cleanDb(b)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		u := createUserBenchmark(
			bson.NewObjectId(),
			"Peter_"+strconv.Itoa(i),
			"Creative",
			b,
		)
		b.StartTimer()
		// Delete Record Test
		err := Delete(u.ID)
		if err != nil {
			b.Fatalf("Error removing user record: %s", err)
		}
	}
}

func BenchmarkCRUD(b *testing.B) {
	os.Remove(dbPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

		// Create Record Test
		u1 := createUserBenchmark(
			bson.NewObjectId(),
			"Peter",
			"Creative",
			b,
		)

		// Read Record Test
		_, err := One(u1.ID)
		if err != nil {
			b.Fatalf("Error retrieving an existing user: %s", err)
		}

		// Update Record Test
		u1.Role = "Documents engineering"
		err = u1.Save()
		if err != nil {
			b.Fatalf("Error saving a record: %s", err)
		}

		// Delete Record Test
		err = Delete(u1.ID)
		if err != nil {
			b.Fatalf("Error removing user record: %s", err)
		}
	}
}
