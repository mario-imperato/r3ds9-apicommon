package linkedservices

import (
	"context"
	"errors"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-aws-common/s3/awss3lks"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-mongo-common/mongolks"
	kafka_go "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/mario-imperato/r3ds9-apicommon/linkedservices/kafka"
	"github.com/mario-imperato/r3ds9-apicommon/linkedservices/restclient"
	"github.com/rs/zerolog/log"
)

type ServiceRegistry struct {
	RestClient *restclient.LinkedService
	kafka      []*kafka.LinkedService
}

var registry ServiceRegistry

func InitRegistry(cfg *Config) error {

	registry = ServiceRegistry{}
	log.Info().Msg("initialize services registry")

	_, err := mongolks.Initialize(cfg.Mongo)
	if err != nil {
		return err
	}

	_, err = awss3lks.Initialize(cfg.S3)
	if err != nil {
		return err
	}

	err = initializeRestClientProvider(cfg.RestClient)
	if err != nil {
		return err
	}

	err = initializeKafka(cfg.Kafka)
	if err != nil {
		return err
	}

	return nil
}

/*
 * RestClient Initialization
 */

func initializeRestClientProvider(cfg *restclient.Config) error {
	log.Info().Msg("initializing rest-client-provider")
	if cfg != nil {
		lks, err := restclient.NewInstanceWithConfig(cfg)
		if err != nil {
			return err
		}

		registry.RestClient = lks
	}

	return nil
}

func GetRestClientProvider(opts ...restclient.Option) (*restclient.Client, error) {
	if registry.RestClient != nil {
		return registry.RestClient.NewClient(opts...)
	}

	return nil, errors.New("rest client linked service not available")
}

/*
 * Kafka Initialization
 */

func initializeKafka(cfg []kafka.Config) error {
	log.Info().Msg("initializing kafka backend service")

	for _, kcfg := range cfg {
		if cfg != nil {
			azk, err := kafka.NewKafkaServiceInstanceWithConfig(kcfg)
			if err != nil {
				return err
			}

			registry.kafka = append(registry.kafka, azk)
			log.Info().Msg("kafka instance configured")
		}
	}
	return nil
}

func GetKafkaLinkedService(brokerName string) (*kafka.LinkedService, error) {

	log.Trace().Str("broker", brokerName).Msg("get kafka linked service")

	for _, k := range registry.kafka {
		if k.Name() == brokerName {
			return k, nil
		}
	}

	err := errors.New("kafka linked service not configured")
	log.Error().Err(err).Str("broker-name", brokerName).Msg("retrieve kafka linked service")
	return nil, err
}

func NewKafkaConsumer(brokerName, gId string) (*kafka_go.Consumer, error) {

	k, err := GetKafkaLinkedService(brokerName)
	if err != nil {
		return nil, err
	}

	return k.NewConsumer(gId)
}

func NewKafkaProducer(ctx context.Context, brokerName, tId string) (*kafka_go.Producer, error) {

	k, err := GetKafkaLinkedService(brokerName)
	if err != nil {
		return nil, err
	}

	return k.NewProducer(ctx, tId)
}

/*
 *


const MongoDbDefaultInstanceName = "default"

func GetMongoDbService(ctx context.Context, n string) (*mongodb.MDbLinkedService, error) {
	log.Trace().Str("broker", n).Msg("get mongo linked service")

	for _, k := range registry.MDbLinkedService {
		if k.Name() == n {
			if !k.IsConnected() {
				err := k.Connect(ctx)
				if err != nil {
					return nil, err
				}
			}
			return k, nil
		}
	}

	err := errors.New("mongo linked service not configured")
	log.Error().Err(err).Str("broker-name", n).Msg("retrieve mongo linked service")
	return nil, err
}

func GetMongoDbCollection(ctx context.Context, instanceName string, collectionId string) (*mongo.Collection, error) {

	lks, err := GetMongoDbService(ctx, instanceName)
	if err != nil {
		return nil, err
	}

	c := lks.GetCollection(collectionId, "")
	if c == nil {
		return c, fmt.Errorf("cannot find collection by id %s", collectionId)
	}

	return c, nil
}

func initializeMongo(cfg []mongodb.MongoConfig) error {
	log.Info().Msg("initializing Mongodb backend service")

	for _, kcfg := range cfg {
		if cfg != nil {
			azk, err := mongodb.NewLinkedService(kcfg)
			if err != nil {
				return err
			}

			registry.MDbLinkedService = append(registry.MDbLinkedService, azk)
			log.Info().Msg("mongodb instance configured")
		}
	}
	return nil
}
*/
