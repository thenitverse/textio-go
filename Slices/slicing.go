package main

import (
	"errors"
	"reflect"
	"testing"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == "pro" {
		return messages[:], nil
	} else if plan == "free" {
		return messages[:2], nil
	}
	return nil, errors.New("unsupported plan")
}

func TestGetMessageWithRetriesForPlan(t *testing.T) {
	messages := [3]string{"msg1", "msg2", "msg3"}

	t.Run("Pro Plan", func(t *testing.T) {
		got, err := getMessageWithRetriesForPlan("pro", messages)
		want := []string{"msg1", "msg2", "msg3"}
		if err != nil {
			t.Errorf("got error %v, want nil", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Free Plan", func(t *testing.T) {
		got, err := getMessageWithRetriesForPlan("free", messages)
		want := []string{"msg1", "msg2"}
		if err != nil {
			t.Errorf("got error %v, want nil", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Unsupported Plan", func(t *testing.T) {
		got, err := getMessageWithRetriesForPlan("business", messages)
		if got != nil {
			t.Errorf("got %v, want nil", got)
		}
		if err == nil || err.Error() != "unsupported plan" {
			t.Errorf("got unexpected error: %v", err)
		}
	})
}
