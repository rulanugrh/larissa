package pkg

import (
	"github.com/prometheus/client_golang/prometheus"
)

type IPrometheus interface {
	SetGauge() *Data
}

type Data struct {
	Kunjungan        prometheus.Gauge
	Obat             prometheus.Gauge
	User             prometheus.Gauge
	Penyakit         prometheus.Gauge
	KunjunganUpgrade *prometheus.CounterVec
	ObatUpgrade      *prometheus.CounterVec
	UserUpgrade      *prometheus.CounterVec
	PenyakitUpgrade  *prometheus.CounterVec
	KunjunganHistory *prometheus.HistogramVec
	ObatHistory      *prometheus.HistogramVec
	UserHistory      *prometheus.HistogramVec
	PenyakitHistory  *prometheus.HistogramVec
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
			Namespace: "kunjungan",
			Name:      "list_all_kunjungan",
			Help:      "Banyaknya user yang telah melakukan kunjungan",
		}),
		Obat: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "obat",
			Name:      "list_all_obat",
			Help:      "Banyaknya user yang telah melihat list obat",
		}),
		User: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "user",
			Name:      "list_all_user",
			Help:      "Banyaknya user yang telah melakukan register",
		}),
		Penyakit: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "penyakit",
			Name:      "list_all_penyakit",
			Help:      "Banyaknya user yang telah melihat list daripada penyakit",
		}),
		PenyakitUpgrade: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "penyakit",
			Name:      "penyakit_update_counter",
			Help:      "For create, update, and deleted",
		}, []string{"type"}),
		ObatUpgrade: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "obat",
			Name:      "obat_update_counter",
			Help:      "For create, update, and deleted",
		}, []string{"type"}),
		UserUpgrade: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "user",
			Name:      "user_update_counter",
			Help:      "For create, update, and deleted",
		}, []string{"type"}),
		KunjunganUpgrade: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "kunjungan",
			Name:      "kunjungan_update_counter",
			Help:      "For create, update, and deleted",
		}, []string{"type"}),
		KunjunganHistory: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "kunjungan",
			Name: "kunjungan_histogram",
			Help: "Histogram for any request in endpoint kunjungan",
		}, []string{"method", "code", "type"}),
		ObatHistory: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "obat",
			Name: "obat_histogram",
			Help: "Histogram for any request in endpoint obat",
		}, []string{"method", "code", "type"}),
		UserHistory: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "user",
			Name: "user_histogram",
			Help: "Histogram for any request in endpoint user",
		}, []string{"method", "code", "type"}),
		PenyakitHistory: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "penyakit",
			Name: "penyakit_histogram",
			Help: "Histogram for any request in endpoint penyakit",
		}, []string{"method", "code", "type"}),
	}

	m.reg.MustRegister(ms.Kunjungan, ms.Obat, ms.Penyakit, ms.User, ms.KunjunganUpgrade, ms.ObatUpgrade, ms.PenyakitUpgrade, ms.UserUpgrade, ms.KunjunganHistory, ms.ObatHistory, ms.UserHistory, ms.PenyakitHistory)

	return &ms
}
