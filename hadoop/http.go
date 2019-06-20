//	hadoop.http 封装了http请求，所有的exporter可以引用

package hadoop

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/common/log"
)

// JmxEnvelope Jmx metrics 对应转化的结构
type JmxEnvelope struct {
	Beans []map[string]interface{} `json:"beans"`
}

// GetJmxEnvelope 通过http 获取Jmx的信息
func GetJmxEnvelope() (*JmxEnvelope, error) {
	httpClient := http.Client{}
	resp, err := httpClient.Get(JmxURL)
	if err != nil {

		log.Errorf("Failed to collect metrics from %v: %s", JmxURL, err)
		return nil, err
	}
	defer func() {
		ioutil.ReadAll(resp.Body) // Mindless drain body upon exit
		resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Failed to collect metrics from %v: HTTP status code %d", JmxURL, resp.StatusCode)
		return nil, fmt.Errorf("Failed to collect metrics from %v: HTTP status code %d", JmxURL, resp.StatusCode)
	}

	var envelope JmxEnvelope
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&envelope)
	if err != nil {
		log.Errorf("Failed to collect metrics from %v: %s", JmxURL, err)
		return nil, err
	}
	return &envelope, nil
}
