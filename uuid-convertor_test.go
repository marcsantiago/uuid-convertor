package convertor

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestUUIDToIntString(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should consistantly return the same string representation",
			args: args{
				uuid: "8fa7158f-f262-44c2-8633-f02298624c52",
			},
			want: "190947154307861374080435443858447420498",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UUIDToIntString(tt.args.uuid); got != tt.want {
				t.Errorf("UUIDToIntString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUUIDToUInt(t *testing.T) {
	type args struct {
		uuid string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "should consistantly return the same string representation",
			args: args{
				uuid: "8fa7158f-f262-44c2-8633-f02298624c52",
			},
			want: 9670336856270720082,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UUIDToUInt(tt.args.uuid); got != tt.want {
				t.Errorf("UUIDToUInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestUUIDToIntStringCollision(t *testing.T) {
	idMap := make(map[string]struct{})
	for i := 0; i < 10000000; i++ {
		id, _ := uuid.NewV4()
		nid := UUIDToIntString(id.String())
		if _, ok := idMap[nid]; ok {
			t.Fatal("collision was detected when coverting to a string")
		}
	}
}

func TestUUIDToUIntCollision(t *testing.T) {
	idMap := make(map[uint64]struct{})
	for i := 0; i < 10000000; i++ {
		id, _ := uuid.NewV4()
		nid := UUIDToUInt(id.String())
		if _, ok := idMap[nid]; ok {
			t.Fatal("collision was detected when coverting to a uint64")
		}
	}
}

func BenchmarkUUIDToIntString(b *testing.B) {
	uid := "8fa7158f-f262-44c2-8633-f02298624c52"
	for i := 0; i < b.N; i++ {
		UUIDToIntString(uid)
	}
}

func BenchmarkUUIDToUInt(b *testing.B) {
	uid := "8fa7158f-f262-44c2-8633-f02298624c52"
	for i := 0; i < b.N; i++ {
		UUIDToUInt(uid)
	}
}
