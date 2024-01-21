package worker

import (
	"github.com/surendratiwari3/paota/broker"
	amqpBroker "github.com/surendratiwari3/paota/broker/amqp"
	brokerErrors "github.com/surendratiwari3/paota/broker/errors"
	"github.com/surendratiwari3/paota/config"
	"github.com/surendratiwari3/paota/store"
	storeErrors "github.com/surendratiwari3/paota/store/errors"
	nullStore "github.com/surendratiwari3/paota/store/null"
)

// CreateBroker creates a new object of broker.Broker
func CreateBroker(cnf *config.Config) (broker.Broker, error) {
	switch cnf.Broker {
	case "amqp":
		return amqpBroker.NewAMQPBroker(cnf)
	default:
		return nil, brokerErrors.ErrUnsupportedBroker
	}
}

// CreateStore creates a new object of store.Interface
func CreateStore(cnf *config.Config) (store.Backend, error) {
	switch cnf.Store {
	case "null":
		return nullStore.NewNullBackend(), nil
	default:
		return nil, storeErrors.ErrUnsupportedStore
	}
}