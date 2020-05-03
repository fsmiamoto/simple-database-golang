package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	EXIT_SUCCESS = 0
)

type metaCommandExecStatus int

const (
	SUCESS metaCommandExecStatus = iota
	UNRECOGNIZED
)

type Statement struct {
	stype statementType
}

type statementType int

const (
	INSERT statementType = iota
	SELECT
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		printPrompt()

		if ok := scanner.Scan(); !ok {
			break
		}
		cmd := getCommandFromScanner(scanner)

		if isMetaCommand(cmd) {
			err := execMetaCommand(cmd)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		stmt, err := buildStatementFromCommand(cmd)
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = execStatment(stmt)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func printPrompt() {
	fmt.Print("db > ")
}

func getCommandFromScanner(s *bufio.Scanner) string {
	rawCmd := s.Text()
	trimmed := strings.TrimSpace(rawCmd)
	return strings.ToLower(trimmed)
}

func execMetaCommand(input string) error {
	if input == ".exit" {
		os.Exit(EXIT_SUCCESS)
	}
	return errors.New("unrecognized meta-command")
}

func isMetaCommand(command string) bool {
	return strings.HasPrefix(command, ".")
}

func buildStatementFromCommand(command string) (*Statement, error) {
	stmt := &Statement{}

	if strings.HasPrefix(command, "select") {
		stmt.stype = SELECT
		return stmt, nil
	}

	if strings.HasPrefix(command, "insert") {
		stmt.stype = INSERT
		return stmt, nil
	}

	return nil, errors.New("unrecognized statement")
}

func execStatment(stmt *Statement) error {
	switch stmt.stype {
	case INSERT:
		fmt.Println("This is where we would make an insert")
	case SELECT:
		fmt.Println("This is where we would make an select")
	}
	return nil
}
