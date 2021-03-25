package ugener

import "testing"

func TestSnowflake(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := RandomWorker().GetId()
		output := Short63(id)
		t.Log(output)
	}

}

func TestHexSnowflake(t *testing.T) {
	for i := 0; i < 10; i++ {
		worker := RandomWorker()
		t.Log(worker.WorkerId(), RandomWorker().GetStringHexTrim())
	}
}
