package models

type ServiceInfo struct {
	Name string
	Desc string
	URL  string
	Up   bool
}

func GetDefaultServicesInfos() []ServiceInfo {
	return []ServiceInfo{
		{"Frontend", "", "http://tagenal", true},
		{"Users API", "", "http://api.tagenal/users", true},
		{"Articles API", "", "http://api.tagenal/articles", true},
		{"Grafana", "", "http://grafana.tagenal", true},
		{"Prometheus", "", "http://prometheus.tagenal", true},
		{"AlertManager", "", "http://alertmanager.tagenal", true},
		{"Traefik", "", "http://tagenal:8080", true},
		{"Vtctld Dashboard", "", "http://tagenal/vtctld/app/dashboard", true},
	}
}
