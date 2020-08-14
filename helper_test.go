package main

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/assert"
)

func bigEqual(t *testing.T, x, y *big.Int) {
	assert.True(t, x.Cmp(y) == 0, "expected %v, got %v", x, y)
}

func strContains(t *testing.T, container, substr string) {
	assert.True(t, strings.Contains(container, substr), "expected to contain `%v`, got `%v`", substr, container)
}

func skipTx(_ *types.Transaction, err error) error {
	return err
}
