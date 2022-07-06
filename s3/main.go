package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
	"strings"
)

var (
	client *s3.Client
)

func init() {
	log.Println("load config")
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile("zigzag-data"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("create s3 client")
	client = s3.NewFromConfig(cfg)
}

func main() {
	//checkItConfigSvc()
	//changeItConfigSvc()
	//applySameConfig()
	//()
	//getS3VersionedList()
	deleteLogs()
}

func test() {
	bucketName := "rookie-version"
	objectVersionsOutput, err := client.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
		Bucket: &bucketName,
	})
	if err != nil {
		panic(err)
	}
	for _, version := range objectVersionsOutput.Versions {
		if !version.IsLatest {
			//fmt.Println(*version.Key, *version.VersionId)
			tempObject := *version.Key
			if tempObject[len(tempObject)-1] == 47 {
				log.Printf("해당 버전관리 대상은 디렉토리입니다. 버전관리 대상 명: %s versionID: %s \n", tempObject, *version.VersionId)
				continue
			}
			if *version.Key == "_$folder$" {
				log.Printf(" _$folder$ 디렉토리는 삭제하지 않습니다. %s versionID: %s \n", tempObject, *version.VersionId)
				continue
			}

			// 1-2 3개의 디렉토리에서 최근 데이터를 제외하고 버전관리된 데이터를 모두 삭제
			//if strings.Contains(*version.Key, "zigzag/mart/log/user_behavior_logs_partitioned/") ||
			//	strings.Contains(*version.Key, "fbk/mart/log/user_behavior_logs_partitioned/") ||
			//	strings.Contains(*version.Key, "zibet/mart/log/user_behavior_logs_partitioned/") {
			//	log.Printf("%s 는 delete marker를 삭제해 복구한 데이터입니다. 최근데이터 이전의 버전관리 데이터는 모두 삭제합니다. \n", *version.Key)
			//	deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
			//} else {
			// 3개의 디렉토리가 아닌경우 버전 데이터 모두 삭제
			fmt.Printf("[%s] %s 의 버전데이터를 삭제합니다. 버전아이디: %s \n", version.StorageClass, *version.Key, *version.VersionId)
			deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
			//}
		}
	}
}

func deleteLog() {
	bucketName := "croquis-data-emr"
	objectVersionsOutput, err := client.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
		Bucket: &bucketName,
	})
	if err != nil {
		panic(err)
	}

	for _, version := range objectVersionsOutput.Versions {
		if !strings.HasPrefix(*version.Key, "2022-05-02-") {
			break
		}
		log.Printf("delete object: %s", *version.Key)
		deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
		//deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
	}

}

func deleteLogs() {
	//var ListObjectVersionsOutputList []*s3.ListObjectVersionsOutput
	bucketName := "croquis-data-emr"
	keyMarker := ""
	for {
		objectVersionsOutput, err := client.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
			Bucket:    &bucketName,
			KeyMarker: &keyMarker,
		})
		if err != nil {
			panic(err)
		}
		log.Printf("[%d]개 단위로 쪼개서 가져옵니다, [%d]개중 첫번째 object는 %s 입니다. \n", len(objectVersionsOutput.Versions), len(objectVersionsOutput.Versions), *objectVersionsOutput.Versions[0].Key)
		for _, version := range objectVersionsOutput.Versions {
			if strings.HasPrefix(*version.Key, "2022-05-02-") || strings.HasPrefix(*version.Key, "2022-05-03-") {
				log.Printf("delete object: %s", *version.Key)
				deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
			} else {
				log.Println("finish")
			}
			//deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
		}
		if !objectVersionsOutput.IsTruncated {
			break
		} else {
			keyMarker = *objectVersionsOutput.NextKeyMarker
		}
	}
	fmt.Println("finish")
	//fmt.Println(len(ListObjectVersionsOutputList))
}

func executeDeleteVersionedObject() {
	//var ListObjectVersionsOutputList []*s3.ListObjectVersionsOutput
	bucketName := "croquis-data-emr"
	keyMarker := ""
	for {
		objectVersionsOutput, err := client.ListObjectVersions(context.TODO(), &s3.ListObjectVersionsInput{
			Bucket:    &bucketName,
			KeyMarker: &keyMarker,
		})
		if err != nil {
			panic(err)
		}
		//ListObjectVersionsOutputList = append(ListObjectVersionsOutputList, objectVersionsOutput)

		if len(objectVersionsOutput.Versions) == 0 {
			log.Println("delete marker만 뽑았으므로 추가작업을 진행하지 않습니다. 추후 delete marker를 삭제하는 작업이 필요합니다.")
			if !objectVersionsOutput.IsTruncated {
				break
			} else {
				keyMarker = *objectVersionsOutput.NextKeyMarker
			}
			continue
		}
		log.Printf("[%d]개 단위로 쪼개서 가져옵니다, [%d]개중 첫번째 object는 %s 입니다. \n", len(objectVersionsOutput.Versions), len(objectVersionsOutput.Versions), *objectVersionsOutput.Versions[0].Key)

		/**
		version.IsLatest 꼭 확인해야함
		*/
		for _, version := range objectVersionsOutput.Versions {
			if !version.IsLatest {
				//fmt.Println(*version.Key, *version.VersionId)
				tempObject := *version.Key
				if tempObject[len(tempObject)-1] == 47 {
					log.Printf("해당 버전관리 대상은 디렉토리입니다. 버전관리 대상 명: %s versionID: %s \n", tempObject, *version.VersionId)
					continue
				}
				if *version.Key == "_$folder$" {
					log.Printf(" _$folder$ 디렉토리는 삭제하지 않습니다. %s versionID: %s \n", tempObject, *version.VersionId)
					continue
				}

				// 1-2 3개의 디렉토리에서 최근 데이터를 제외하고 버전관리된 데이터를 모두 삭제
				//if strings.Contains(*version.Key, "zigzag/mart/log/user_behavior_logs_partitioned/") ||
				//	strings.Contains(*version.Key, "fbk/mart/log/user_behavior_logs_partitioned/") ||
				//	strings.Contains(*version.Key, "zibet/mart/log/user_behavior_logs_partitioned/") {
				//	log.Printf("%s 는 delete marker를 삭제해 복구한 데이터입니다. 최근데이터 이전의 버전관리 데이터는 모두 삭제합니다. \n", *version.Key)
				//	deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
				//} else {
				// 3개의 디렉토리가 아닌경우 버전 데이터 모두 삭제
				fmt.Printf("[%s] %s 의 버전데이터를 삭제합니다. 버전아이디: %s \n", version.StorageClass, *version.Key, *version.VersionId)
				deleteObjectWithVersionId(bucketName, *version.Key, *version.VersionId)
				//}
			}
		}
		if !objectVersionsOutput.IsTruncated {
			break
		} else {
			keyMarker = *objectVersionsOutput.NextKeyMarker
		}
	}
	fmt.Println("finish")
	//fmt.Println(len(ListObjectVersionsOutputList))
}

func deleteObject(bucketName string, objectName string) {
	_, err := client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: &bucketName,
		Delete: &types.Delete{
			Objects: []types.ObjectIdentifier{
				{
					Key:       &objectName,
					VersionId: nil,
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}

func deleteObjectWithVersionId(bucketName string, objectName string, versionId string) {
	//log.Printf("[%s]버킷의 [%s]버전아이디의 [%s]오브젝트를 삭제합니다. \n", bucketName, versionId, objectName)
	_, err := client.DeleteObjects(context.TODO(), &s3.DeleteObjectsInput{
		Bucket: &bucketName,
		Delete: &types.Delete{
			Objects: []types.ObjectIdentifier{
				{
					Key:       &objectName,
					VersionId: &versionId,
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}

func applySameConfig() {
	// 기존의 config는 aws 콘솔에서 직접 생성한 config
	exBucket := "rookie1"
	exIntelligentTieringId := "it-config"
	configurationOutput, err := client.GetBucketIntelligentTieringConfiguration(context.TODO(), &s3.GetBucketIntelligentTieringConfigurationInput{
		Bucket: &exBucket,
		Id:     &exIntelligentTieringId,
	})
	if err != nil {
		panic(err)
	}
	// 콘솔에 있는 값을 가져오는건 성공
	fmt.Println("success to get intelligent tiering config")

	// bucket2에 해당 config를 적용
	newBucket := ""
	newIntelligentTieringId := "new-config"
	input := &s3.PutBucketIntelligentTieringConfigurationInput{
		Bucket:                          &newBucket,
		Id:                              &newIntelligentTieringId,
		IntelligentTieringConfiguration: configurationOutput.IntelligentTieringConfiguration, // 콘솔에서 만든 값을 이용해서 변경해도 400

		// // https://docs.aws.amazon.com/cli/latest/reference/s3api/put-bucket-intelligent-tiering-configuration.html
		// // 직접 struct를 만들어줘도 400
		//IntelligentTieringConfiguration: &types.IntelligentTieringConfiguration{
		//	Id:     &newIntelligentTieringId,
		//	Status: types.IntelligentTieringStatusEnabled,
		//
		//  // filter를 주지않고 버킷안에 존재하는 모든 object에 대해서 해당 config를 적용하려고 시도
		//	// filter struct를 만들어서 내부에 nil을 줘도 400, And operator의 struct를 만들어서 값에 nil을 줘도 400
		//	//Filter: &types.IntelligentTieringFilter{
		//	//	And:    nil,
		//	//	Prefix: nil,
		//	//	Tag:    nil,
		//	//},
		//
		//	// filter를 nil로 해도, filter 필드를 아에 만들지 않아도 400
		//	Filter: nil,
		//
		//	Tierings: []types.Tiering{
		//		{
		//			AccessTier: types.IntelligentTieringAccessTierArchiveAccess,
		//			Days:       200,
		//		},
		//		{
		//			AccessTier: types.IntelligentTieringAccessTierDeepArchiveAccess,
		//			Days:       100,
		//		},
		//	},
		//},
	}
	_, err = client.PutBucketIntelligentTieringConfiguration(context.TODO(), input)
	if err != nil {
		// 400 error 발생
		panic(err)
	}
	fmt.Println("success to put intelligent tiering config")
}

// 3월 15일부터 버저닝이 두개인거중에 하나는 스탠다드 하나는 인텔리전트 티어링인 object리스트를 뽑아보기

func checkItConfigSvc() {
	bucketName := "rookie1"
	itId := "it-config"
	configuration, err := client.GetBucketIntelligentTieringConfiguration(context.TODO(), &s3.GetBucketIntelligentTieringConfigurationInput{
		Bucket: &bucketName,
		Id:     &itId,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("get bucket it config: ", configuration)

	newBucket := "my--test1"
	ITConfigId := "fabc"
	input := &s3.PutBucketIntelligentTieringConfigurationInput{
		Bucket:                          &newBucket,
		Id:                              &ITConfigId,
		IntelligentTieringConfiguration: configuration.IntelligentTieringConfiguration,
	}
	_, err = client.PutBucketIntelligentTieringConfiguration(context.TODO(), input)
	if err != nil {
		panic(err)
	}
}

//func () {
//	// 기존의 config는 aws 콘솔에서 직접 생성한 config
//	exBucket := "my--test1"
//	exIntelligentTieringId := "test-config"
//	configurationOutput, err := client.GetBucketIntelligentTieringConfiguration(context.TODO(), &s3.GetBucketIntelligentTieringConfigurationInput{
//		Bucket: &exBucket,
//		Id:     &exIntelligentTieringId,
//	})
//	if err != nil {
//		panic(err)
//	}
//	// 콘솔에 있는 값을 가져오는건 성공(버킷에 직접 콘솔에서 설정한 값)
//	fmt.Println("success to get intelligent tiering config")
//
//	// bucket2에 해당 config를 적용
//	newBucket := "my--test2"
//	newIntelligentTieringId := "new-config"
//	input := &s3.PutBucketIntelligentTieringConfigurationInput{
//		Bucket:                          &newBucket,
//		Id:                              &newIntelligentTieringId,
//		IntelligentTieringConfiguration: configurationOutput.IntelligentTieringConfiguration, // 콘솔에서 만든 값을 이용해서 변경해도 400
//	}
//	_, err = client.PutBucketIntelligentTieringConfiguration(context.TODO(), input)
//	if err != nil {
//		// 400 error
//		// MalformedXML: The XML you provided was not well-formed or did not validate against our published schema
//		panic(err)
//	}
//	fmt.Println("success to put intelligent tiering config")
//}

func changeItConfigSvc() {
	bucketList, err := getBucketList()
	if err != nil {
		panic(err)
	}
	for _, bucket := range bucketList.Buckets {
		if *bucket.Name != "rookie1" {
			continue
		}
		_, err := changeBucketIntelligentTieringWithFullArchive(bucket)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func changeBucketIntelligentTieringWithFullArchive(bucket types.Bucket) (*s3.PutBucketIntelligentTieringConfigurationOutput, error) {
	ITConfigId := "full-archive-it-configId"
	input := &s3.PutBucketIntelligentTieringConfigurationInput{
		Bucket: bucket.Name,
		Id:     &ITConfigId,
		IntelligentTieringConfiguration: &types.IntelligentTieringConfiguration{
			Id:     &ITConfigId,
			Status: types.IntelligentTieringStatusEnabled,
			// 아래의 경우에도 실패
			//Filter: &types.IntelligentTieringFilter{
			//	And:    nil,
			//	Prefix: nil,
			//	Tag:    nil,
			//},
			Filter: nil,
			Tierings: []types.Tiering{
				{
					AccessTier: types.IntelligentTieringAccessTierArchiveAccess,
					Days:       200,
				},
				{
					AccessTier: types.IntelligentTieringAccessTierDeepArchiveAccess,
					Days:       100,
				},
			},
		},
	}
	fmt.Println("test")
	return client.PutBucketIntelligentTieringConfiguration(context.TODO(), input)
}

func getBucketList() (*s3.ListBucketsOutput, error) {
	buckets, err := client.ListBuckets(context.TODO(), nil)
	if err != nil {
		log.Println("버킷을 가져오는데 실패했습니다.", err)
		return nil, err
	}
	return buckets, nil
}

func getBucketInfo(bucketName string) {
}
