package utils

import (
	"os"
	"testing"
)

func TestConfigReadingFromOS(t *testing.T) {
	config := NewConfig()

	// before setting env variable
	if v := config.Get("config_test"); v != "" {
		t.Fatalf("found config_test env variable with value '%s'", v)
	}

	// Now set it
	os.Setenv("config_test", "hi")
	if v := config.Get("config_test"); v != "hi" {
		t.Fatalf("config_test env variable should have value hi, but has '%s'", v)
	}

	os.Unsetenv("config_test")
}

func TestConfigReadingFromFile(t *testing.T) {
	config := NewConfigFromPath("../testdata/horizon-config")

	if v := config.Get("empty_in_file"); v != "" {
		t.Fatalf("expected empty config value for empty_in_file got '%s'", v)
	}

	// blanked is not an override
	os.Setenv("empty_in_file", "some-value")
	if v := config.Get("empty_in_file"); v != "some-value" {
		t.Fatalf("expected empty_in_file to be some-value got '%s'", v)
	}

	if v := config.Get("exists_only_in_file"); v != "A" {
		t.Fatalf("expected exists_only_in_file equals A, got '%s'", v)
	}

	// override from os.Getenv
	if v := config.Get("exists_in_os_env_and_file"); v != "B" {
		t.Fatalf("expected exists_in_os_env_and_file equals B, got '%s'", v)
	}
}

func TestConfigFileParsingHandlesValuesWithEquals(t *testing.T) {
	config := NewConfigFromPath("../testdata/horizon-config")

	if v := config.Get("contains_multiple_equals"); v != "like_some=passwords_would" {
		t.Fatalf("expected contains_multiple_equals equals like_some=passwords_would, got '%s'", v)
	}
}
