package main

import "fmt"

type ElasticWebSpec struct {
	Image        string `json:"image"`
	Port         *int32 `json:"port"`
	SinglePodQPS *int32 `json:"singlePodQPS"`
	TotalQPS     *int32 `json:"totalQPS"`
}

type ElasticWebStatus struct {
	RealQPS *int32 `json:"realQPS,omitempty"`
}

type ElasticWeb struct {
	//metav1.TypeMeta   `json:",inline"`
	//metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec   ElasticWebSpec   `json:"spec,omitempty"`
	Status ElasticWebStatus `json:"status,omitempty"`
}

func main() {
	port := int32(3000)
	sqps := int32(500)
	tqps := int32(2000)
	realQPS := int32(600)

	a := &ElasticWeb{
		Spec: ElasticWebSpec{
			Image:        "xxxxx:1.0",
			Port:         &port,
			SinglePodQPS: &sqps,
			TotalQPS:     &tqps,
		},
		Status: ElasticWebStatus{RealQPS: &realQPS},
	}
	fmt.Println(a)
}
