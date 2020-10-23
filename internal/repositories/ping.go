package repositories

import (
	"github.com/NodeFactoryIo/vedran/internal/models"
	"github.com/asdine/storm/v3"
)

type PingRepository interface {
	FindByNodeID(nodeId string) (*models.Ping, error)
	Save(ping *models.Ping) error
	GetAll() (*[]models.Ping, error)
}

type pingRepo struct {
	db *storm.DB
}

func NewPingRepo(db *storm.DB) PingRepository {
	return &pingRepo{
		db: db,
	}
}

func (r *pingRepo) FindByNodeID(nodeId string) (*models.Ping, error) {
	var ping models.Ping
	err := r.db.One("NodeId", nodeId, &ping)
	return &ping, err
}

func (r *pingRepo) Save(ping *models.Ping) error {
	return r.db.Save(ping)
}

func (r *pingRepo) GetAll() (*[]models.Ping, error) {
	var pings []models.Ping
	err := r.db.All(&pings)
	return &pings, err
}


