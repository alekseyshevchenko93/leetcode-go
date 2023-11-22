package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProblem1(t *testing.T) {
	node1 := ListNode{Val: 1}
	node2 := ListNode{Val: 2}
	node3 := ListNode{Val: 3}

	node1.Next = &node2
	node2.Next = &node3

	require.Equal(t, &node3, reverseList(&node1))
	fmt.Println(node3.Next.Val)
	require.Equal(t, node3.Next, &node2)
	require.Equal(t, node2.Next, &node1)
	require.Equal(t, node1.Next, (*ListNode)(nil))
}
