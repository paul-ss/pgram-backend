package single_transaction

import "fmt"

type SavePoint interface {
	Create(id string) string
	Release(id string) string
	Rollback(id string) string
}

type defaultSavePoint struct{}

func (dsp *defaultSavePoint) Create(id string) string {
	return fmt.Sprintf("SAVEPOINT %s", id)
}
func (dsp *defaultSavePoint) Release(id string) string {
	return fmt.Sprintf("RELEASE SAVEPOINT %s", id)
}
func (dsp *defaultSavePoint) Rollback(id string) string {
	return fmt.Sprintf("ROLLBACK TO SAVEPOINT %s", id)
}

func SavePointOption(savePoint SavePoint) func(*conn) error {
	return func(c *conn) error {
		c.savePoint = savePoint
		return nil
	}
}
