package pkg

import "testing"

func TestHash(t *testing.T) {
	type args struct {
		plain string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Hash(tt.args.plain)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Hash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckHash(t *testing.T) {
	type args struct {
		plain  string
		hashed string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckHash(tt.args.plain, tt.args.hashed); got != tt.want {
				t.Errorf("CheckHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
