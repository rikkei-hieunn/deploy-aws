/*
Package model define all constances
*/
package model

const (
	//EmptyString empty string character
	EmptyString = ""
	//Stroke character
	Stroke = "/"
	// MachineSuffix suffix machine id for get instance id from os environment
	MachineSuffix = "_ID"
	//TypeRunAll run all instance
	TypeRunAll ="1"
	//TypeRunSSS run specific instance with group and data type
	TypeRunSSS ="2"
	//TypeRunByGroupLine run specific instance with group line
	TypeRunByGroupLine ="3"

	//FirstKei define first db kei
	FirstKei ="1"
	//SecondKei define second db kei
	SecondKei ="2"
    //TypeTick type for data tick
	TypeTick = "1"
	//TypeKehai type for data kehai
	TypeKehai = "2"
)
const (
	// S3BucketKey key environment save bucket name
	S3BucketKey = "TK_SYSTEM_BUCKET_NAME"
	// S3RegionKey key environment save region
	S3RegionKey = "TK_SYSTEM_REGION"
)