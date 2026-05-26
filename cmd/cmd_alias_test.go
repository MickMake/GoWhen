package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/cobra"
)

func TestLoadAliasesMissingFileReturnsEmptyMap(t *testing.T) {
	cs := aliasTestCmds(t)

	aliases, err := cs.LoadAliases()
	if err != nil {
		t.Fatal(err)
	}
	if len(aliases) != 0 {
		t.Fatalf("LoadAliases len = %d, want 0", len(aliases))
	}
}

func TestSaveAndLoadAliases(t *testing.T) {
	cs := aliasTestCmds(t)

	want := AliasMap{
		"christmas": {"parse", ".", "2026-12-25"},
		"workday":   {"drop", "weekend"},
	}

	if err := cs.SaveAliases(want); err != nil {
		t.Fatal(err)
	}

	got, err := cs.LoadAliases()
	if err != nil {
		t.Fatal(err)
	}

	assertAliasChain(t, got, "christmas", want["christmas"])
	assertAliasChain(t, got, "workday", want["workday"])
}

func TestLoadAliasesInvalidJSONReturnsError(t *testing.T) {
	cs := aliasTestCmds(t)

	if err := os.MkdirAll(filepath.Dir(cs.AliasFile()), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(cs.AliasFile(), []byte("not json"), 0600); err != nil {
		t.Fatal(err)
	}

	if _, err := cs.LoadAliases(); err == nil {
		t.Fatal("expected invalid JSON error")
	}
}

func TestAttachRuntimeAliasValidatesInput(t *testing.T) {
	cs := aliasTestCmds(t)
	root := cobra.Command{Use: "GoWhen"}
	cs.Alias.cmd = &root

	if err := cs.AttachRuntimeAlias(nil, "x", []string{"format", "."}); err == nil {
		t.Fatal("expected nil root error")
	}
	if err := cs.AttachRuntimeAlias(&root, "", []string{"format", "."}); err == nil {
		t.Fatal("expected empty alias name error")
	}
	if err := cs.AttachRuntimeAlias(&root, "empty", nil); err == nil {
		t.Fatal("expected empty chain error")
	}
}

func TestAttachRuntimeAliasRejectsExistingCommand(t *testing.T) {
	cs := aliasTestCmds(t)
	root := cobra.Command{Use: "GoWhen"}
	root.AddCommand(&cobra.Command{Use: "format"})
	cs.Alias.cmd = &root

	if err := cs.AttachRuntimeAlias(&root, "format", []string{"parse", ".", "today"}); err == nil {
		t.Fatal("expected existing command conflict")
	}
}

func TestAttachRuntimeAliasAddsCommand(t *testing.T) {
	cs := aliasTestCmds(t)
	root := cobra.Command{Use: "GoWhen"}
	cs.Alias.cmd = &root

	if err := cs.AttachRuntimeAlias(&root, "christmas", []string{"parse", ".", "2026-12-25"}); err != nil {
		t.Fatal(err)
	}

	if cmd, _, err := root.Find([]string{"christmas"}); err != nil || cmd == nil || cmd.Name() != "christmas" {
		t.Fatalf("expected christmas alias command, cmd=%v err=%v", cmd, err)
	}
}

func TestAttachStoredAliasesAddsCommandsInFreshRoot(t *testing.T) {
	cs := aliasTestCmds(t)
	root := cobra.Command{Use: "GoWhen"}
	cs.Alias.cmd = &root

	aliases := AliasMap{
		"christmas": {"parse", ".", "2026-12-25"},
		"workday":   {"drop", "weekend"},
	}
	if err := cs.SaveAliases(aliases); err != nil {
		t.Fatal(err)
	}

	if err := cs.AttachStoredAliases(&root); err != nil {
		t.Fatal(err)
	}

	for name := range aliases {
		cmd, _, err := root.Find([]string{name})
		if err != nil || cmd == nil || cmd.Name() != name {
			t.Fatalf("expected alias command %q, cmd=%v err=%v", name, cmd, err)
		}
	}
}

func aliasTestCmds(t *testing.T) *Cmds {
	t.Helper()

	cs := &Cmds{
		Unify: cmds.Unify,
		Alias: NewCmdAlias(),
	}

	oldDir := cs.Unify.Commands.CmdConfig.Dir
	cs.Unify.Commands.CmdConfig.Dir = t.TempDir()
	t.Cleanup(func() {
		cs.Unify.Commands.CmdConfig.Dir = oldDir
	})

	return cs
}

func assertAliasChain(t *testing.T, aliases AliasMap, name string, want []string) {
	t.Helper()

	got, ok := aliases[name]
	if !ok {
		t.Fatalf("missing alias %q", name)
	}
	if len(got) != len(want) {
		t.Fatalf("alias %q len = %d, want %d", name, len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("alias %q[%d] = %q, want %q", name, i, got[i], want[i])
		}
	}
}
