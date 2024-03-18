package pkg

import (
	"github.com/prometheus/client_golang/prometheus"
)

type IPrometheus interface {
	SetGauge() *Data
}

type Data struct {
	Kunjungan prometheus.Gauge
	Obat prometheus.Gauge
	User prometheus.Gauge
	Penyakit prometheus.Gauge
}

type Metric struct {
	reg prometheus.Registerer
}

func NewMetric(reg prometheus.Registerer) IPrometheus {
	return &Metric{
		reg: reg,
	}
}

func (m *Metric) SetGauge() *Data {
	ms := Data{
		Kunjungan: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "Kunjungan Namespace",
			Name:      "list_all_kunjungan",
			Help:      "Banyaknya user yang telah melakukan kunjungan",
		}),
		Obat: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "Obat Namespace",
			Name:      "list_all_obat",
			Help:      "Banyaknya user yang telah melihat list obat",
		}),
		User: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "User Namespace",
			Name:      "list_all_user",
			Help:      "Banyaknya user yang telah melakukan register",
		}),
		Penyakit: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "Penyakit Namespace",
			Name:      "list_all_penyakit",
			Help:      "Banyaknya user yang telah melihat list daripada penyakit",
		}),
	}

	m.reg.MustRegister(ms.Kunjungan, ms.Obat, ms.Penyakit, ms.User)

	return &ms
}
