package cache

import (
	"log"
	"testing"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"gotest.tools/assert"
)

func init() {
	InitRedisCache(goredislib.Options{
		Addr:     "localhost:6379",
		Password: "123456",
	})
	setTestValues()
}

func setTestValues() {
	timeout := 3 * time.Second
	type args struct {
		key   string
		value interface{}
	}
	tests := []args{
		{
			key:   "test_string",
			value: "test",
		},
		{
			key:   "test_int",
			value: 5,
		},
		{
			key:   "test_float",
			value: 5.5,
		},
		{
			key:   "test_bool",
			value: true,
		},
		{
			key:   "test_struct",
			value: struct{ Name string }{Name: "test"},
		},
		{
			key:   "test_increase",
			value: 5,
		},
		{
			key:   "test_decrease",
			value: 5,
		},
	}
	for _, tt := range tests {
		if err := cache.Set(tt.key, tt.value, timeout); err != nil {
			log.Fatalf("Set() error = %s", err.Error())
		}
	}
}

func TestDelete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cache.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{"test_string"},
			want:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.GetString(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBool(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{"test_bool"},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.GetBool(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBool() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBool() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFloat64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{"test_float"},
			want:    5.5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.GetFloat64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFloat64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{"test_int"},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.GetInt(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetInt64(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{"test_int"},
			want:    5,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.GetInt64(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInt64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStructData(t *testing.T) {
	type args struct {
		key  string
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				key:  "test_struct",
				data: &struct{ Name string }{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := cache.GetStructData(tt.args.key, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("GetStructData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMutex(t *testing.T) {
	redis, ok := cache.(*Redis)
	if !ok {
		t.Error("redis type covert failed")
	}
	key := "test_mutex"
	count, times := 0, 1000
	for i := 0; i < times; i++ {
		go func(c *int) {
			mutex := redis.GetMutex(key)
			mutex.Lock()
			defer mutex.Unlock()
			*c++
		}(&count)
	}

	wait := make(chan bool)
	go func() {
		time.Sleep(60 * time.Second)
		t.Error("timeout")
		wait <- false
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			if count == times {
				wait <- true
			}
		}
	}()

	assert.Equal(t, <-wait, true)
}

func TestIncrease(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{"test_increase"},
			want:    6,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.Increase(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Increase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Increase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecrease(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "simple",
			args:    args{"test_decrease"},
			want:    4,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := cache.Decrease(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrease() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrease() = %v, want %v", got, tt.want)
			}
		})
	}
}
