package expr

import (
	"fmt"
	"testing"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

type Program struct {
	prog *vm.Program
}

func (prog *Program) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var a string
	if err := unmarshal(&a); err != nil {
		return fmt.Errorf("expression must be a string: %w", err)
	}
	p, err := expr.Compile(a)
	if err != nil {
		return fmt.Errorf("compile a program: %w", err)
	}
	prog.prog = p
	return nil
}

func NewProgram(s string) (*Program, error) {
	prog, err := expr.Compile(s, expr.AsBool())
	if err != nil {
		return &Program{}, fmt.Errorf("compile a program: %w", err)
	}
	return &Program{prog: prog}, nil
}

func NewProgramForTest(t *testing.T, s string) *Program {
	t.Helper()
	a, err := NewProgram(s)
	if err != nil {
		t.Fatal(err)
	}
	return a
}

func (prog *Program) Empty() bool {
	return prog == nil || prog.prog == nil
}

func (prog *Program) Run(param interface{}) (interface{}, error) {
	a, err := expr.Run(prog.prog, param)
	if err != nil {
		return false, fmt.Errorf("evaluate a expr's compiled program: %w", err)
	}
	return a, nil
}
