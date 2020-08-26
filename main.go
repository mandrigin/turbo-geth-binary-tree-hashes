package main

import (
	"fmt"

	"github.com/ledgerwatch/turbo-geth/eth/stagedsync/stages"
)

func main() {
	s := stagedsync.DefaultStages
	s = replaceStage(stages.IntermediateHashes, binaryIntermediateHashes)
	core := turbo.NewCore(
		stagedsync.New(s)
	)
	core.Run()
	fmt.Println("hallo")
}

func replaceStage(source []stagedsync.Stage, id stages.Stage, stage stagedsync.Stages) []stagedsync.Stage {
	panic("implement me")
}
