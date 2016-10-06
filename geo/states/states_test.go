package states

import (
	"testing"
	"github.com/adamdecaf/horizon/utils"
)

func TestReadWriteState(t *testing.T) {
	name := utils.RandString(20)
	empty, err := SearchStatesByName(name)

	if err != nil {
		t.Fatalf("error reading state when we expected to see no results = %s", err)
	}

	if len(empty) != 0 {
		t.Fatal("found states some how when we didn't expect to find any")
	}

	id := utils.RandString(20)
	abbreviation := utils.RandString(2)
	state := State{id, name, abbreviation}

	if written := WriteState(state); written != nil {
		t.Fatalf("error when writing state name=%s, err=%s", name, *written)
	}

	states, err := SearchStatesByName(name)

	if err != nil {
		t.Fatalf("error finding state that should exist name=%s, err=%s\n", name, err)
	}

	if len(states) == 1 {
		found := states[0]
		if found.Id != state.Id || found.Name != state.Name {
			t.Fatalf("states don't match (written=%s) (found=%s)", state, found)
		}
	} else {
		t.Fatalf("found multiple states when we expected one name=%s", name)
	}

	all_states, err := ReadAllStates()
	if err != nil {
		t.Fatalf("not finding any states due to error (err=%s)\n", err)
	}

	if len(all_states) == 0 {
		t.Fatalf("unable to find any states...")
	}

	single_state, err := FindStateById(state.Id)
	if err != nil {
		t.Fatalf("didn't find single state because of error (err=%s)\n", err)
	}

	if single_state == nil {
		t.Fatalf("unable to find single state... (state_id = %s)", state.Id)
	}
}
