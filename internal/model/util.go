package model

import (
	repositoryModel "github.com/FadhilAF/perbakin-be/internal/repository/postgres/sqlc"
)

type StageList int

const (
	Stage0Type StageList = iota
	Stage1Type
	Stage2Type
	Stage3Type
	Stage4Type
	Stage5Type
	Stage6Type
)

const (
	Stage0EndStatus      repositoryModel.Stage0Status = "5"
	Stage0FinishedStatus repositoryModel.Stage0Status = "6"

	Stage13EndStatus      repositoryModel.Stage13Status = "6"
	Stage13FinishedStatus repositoryModel.Stage13Status = "7"

	Stage246EndStatus      repositoryModel.Stage246Status = "3"
	Stage246FinishedStatus repositoryModel.Stage246Status = "4"

	Stage5EndStatus      repositoryModel.Stage5Status = "2"
	Stage5FinishedStatus repositoryModel.Stage5Status = "3"
)

type StageListString string

const (
	Stage0TypeString StageListString = "0"
	Stage1TypeString StageListString = "1"
	Stage2TypeString StageListString = "2"
	Stage3TypeString StageListString = "3"
	Stage4TypeString StageListString = "4"
	Stage5TypeString StageListString = "5"
	Stage6TypeString StageListString = "6"
)
