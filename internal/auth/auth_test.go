package auth

import (
	"testing"
	"time"
)

func TestCheckPasswordHash(t *testing.T) {
	password1 := "RightPassword123"
	password2 := "WrongPassword456"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct Password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Wrong Password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty Password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid Hash",
			password: password1,
			hash:     "invalidHash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	userID := 1
	validToken, _ := MakeJWT(userID, "secret", 30*24*time.Hour)

	tests := []struct {
		name        string
		tokenString string
		tokenSecret string
		wantUserID  int
		wantErr     bool
	}{
		{
			name:        "Valid Token",
			tokenString: validToken,
			tokenSecret: "secret",
			wantUserID:  1,
			wantErr:     false,
		},
		{
			name:        "Invalid Token",
			tokenString: "invalidToken",
			tokenSecret: "secret",
			wantUserID:  0, // 0 is the output from an error in ValidateJWT
			wantErr:     true,
		},
		{
			name:        "WrongToken",
			tokenString: validToken,
			tokenSecret: "wrong_secret",
			wantUserID:  0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(tt.tokenString, tt.tokenSecret)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUserID != tt.wantUserID {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, tt.wantUserID)
			}
		})
	}
}
