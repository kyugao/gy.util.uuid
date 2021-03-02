package ugen

import "testing"

func TestSnowflake(t *testing.T) {
	t.Log(Snowflake.GetString())
}

func TestHexSnowflake(t *testing.T) {
	for i := 0; i < 10; i++ {
		worker := RandomWorker()
		t.Log(worker.WorkerId(), RandomWorker().GetStringHex())
	}
}
