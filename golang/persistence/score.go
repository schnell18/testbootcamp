package persistence

type ScoreDAO interface {
	GetScore(student, subject string) (int, error)
}

func GetAverageScore(db ScoreDAO, student string) (int, error) {
	subjects := []string{"Math", "Chinese", "English"}
	scores := 0
	for _, subject := range subjects {
		if score, err := db.GetScore(student, subject); err != nil {
			return -1, err
		} else {
			scores += score
		}
	}

	return scores / len(subjects), nil
}
