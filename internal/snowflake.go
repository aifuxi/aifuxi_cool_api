package internal

import (
	"github.com/bwmarrin/snowflake"
)

func GenSnowflakeID() (int64, error) {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}

	return node.Generate().Int64(), nil
}
