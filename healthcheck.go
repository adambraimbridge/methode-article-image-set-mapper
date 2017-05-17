package main

import (
	health "github.com/Financial-Times/go-fthealth/v1_1"
	"github.com/Financial-Times/service-status-go/gtg"
	"encoding/json"
	"fmt"
	"errors"
	"github.com/Sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"github.com/Financial-Times/message-queue-gonsumer/consumer"
)

const healthPath = "/__health"

type healthService struct {
	httpClient *http.Client
	consumerConfig consumer.QueueConfig
	config *healthConfig
	checks []health.Check
}

type healthConfig struct {
	appSystemCode string
	appName       string
}

func newHealthService(config *healthConfig, httpClient *http.Client, consumerConfig consumer.QueueConfig) *healthService {
	service := &healthService{
		config: config,
		httpClient: httpClient,
		consumerConfig: consumerConfig,
	}
	service.checks = []health.Check{
		service.messageQueueProxyReachable(),
	}
	return service
}

func (h *healthService) messageQueueProxyReachable() health.Check {
	return health.Check{
		BusinessImpact:   "Publishing or updating image-sets inside article bodies will not be possible, clients will not see images or image-sets in new content.",
		Name:             "MessageQueueProxyReachable",
		PanicGuide:       "https://dewey.ft.com/methode-article-image-set-mapper.html",
		Severity:         1,
		TechnicalSummary: "Message queue proxy is not reachable/healthy",
		Checker:          h.checkAggregateMessageQueueProxiesReachable,
	}
}

func (h *healthService) checkAggregateMessageQueueProxiesReachable() (string, error) {
	errMsg := ""
	for i := 0; i < len(h.consumerConfig.Addrs); i++ {
		err := h.checkMessageQueueProxyReachable(h.consumerConfig.Addrs[i], h.consumerConfig.Topic, h.consumerConfig.AuthorizationKey, h.consumerConfig.Queue)
		if err == nil {
			return "", nil
		}
		errMsg = errMsg + fmt.Sprintf("For %s there is an error %v \n", h.consumerConfig.Addrs[i], err.Error())
	}
	return errMsg, errors.New(errMsg)
}

func (h *healthService) checkMessageQueueProxyReachable(address string, topic string, authKey string, queue string) error {
	req, err := http.NewRequest("GET", address+"/topics", nil)
	if err != nil {
		logrus.Warnf("Could not connect to proxy: %v", err.Error())
		return err
	}
	if len(authKey) > 0 {
		req.Header.Add("Authorization", authKey)
	}
	if len(queue) > 0 {
		req.Host = queue
	}
	resp, err := h.httpClient.Do(req)
	if err != nil {
		logrus.Warnf("Could not connect to proxy: %v", err.Error())
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Proxy returned status: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return checkIfTopicIsPresent(body, topic)
}

func checkIfTopicIsPresent(body []byte, searchedTopic string) error {
	var topics []string
	err := json.Unmarshal(body, &topics)
	if err != nil {
		return fmt.Errorf("Error occured and topic could not be found. %v", err.Error())
	}
	for _, topic := range topics {
		if topic == searchedTopic {
			return nil
		}
	}
	return errors.New("Topic was not found")
}

func (h *healthService) gtgCheck() gtg.Status {
	for _, check := range h.checks {
		if _, err := check.Checker(); err != nil {
			return gtg.Status{GoodToGo: false, Message: err.Error()}
		}
	}
	return gtg.Status{GoodToGo: true}
}
